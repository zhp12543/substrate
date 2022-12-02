package jsonrpc

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func handleResponse(p *WsProvider) error {
	resp := &Response{}
	err := p.client.ReadJSON(resp)
	if err != nil {
		return err
	}

	if resp.Params.Result == nil {
		if c, ok := p.msgChan[resp.ID]; ok {
			c <- resp
		}
	} else {
		id := resp.Params.Subscription
		if callback, ok := p.subscribes[id]; ok {
			if r, ok := p.waitingResponse[id]; ok {
				callback(r.Params.Result)
				delete(p.waitingResponse, id)
			}
			callback(resp.Params.Result)
		} else {
			p.waitingResponse[id] = resp
		}
	}
	return nil
}

func NewWsProvider(endpoint string) (*WsProvider, error) {
	client, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	p := &WsProvider{
		client:          client,
		endpoint:        endpoint,
		id:              0,
		msgChan:         make(map[int]chan *Response),
		subscribes:      make(map[int]func(interface{})),
		waitingResponse: make(map[int]*Response),
		ctx:             ctx,
		ctxCancel:       cancel,
	}

	go func() {
		for {
			select {
			case <-p.ctx.Done():
				return
			default:
				err := handleResponse(p)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}()

	return p, nil
}

type WsProvider struct {
	client          *websocket.Conn
	endpoint        string
	id              int
	msgChan         map[int]chan *Response
	subscribes      map[int]func(interface{})
	waitingResponse map[int]*Response
	ctx             context.Context
	ctxCancel       context.CancelFunc
}

func (p *WsProvider) Call(method string, params []interface{}) (*Response, error) {
	p.id++
	req := &Request{
		ID:      p.id,
		Method:  method,
		JSONRPC: "2.0",
		Params:  params,
	}
	p.msgChan[p.id] = make(chan *Response, 1)
	err := p.client.WriteJSON(req)
	if err != nil {
		return nil, err
	}

	resp := <-p.msgChan[p.id]
	delete(p.msgChan, p.id)
	return resp, nil
}

func (p *WsProvider) Subscribe(method string, params []interface{}, callback func(interface{})) (int, error) {
	resp, err := p.Call(method, params)
	if err != nil {
		return -1, err
	}
	id := int(resp.Result.(float64))
	p.subscribes[id] = callback
	return id, nil
}

func (p *WsProvider) Unsubscribe(method string, id int) error {
	resp, err := p.Call(method, []interface{}{id})
	if err != nil {
		return err
	}
	if resp.Result.(bool) {
		delete(p.subscribes, id)
	}
	return nil
}

func (p *WsProvider) Close() {
	p.ctxCancel()
	p.msgChan = make(map[int]chan *Response)
	p.subscribes = make(map[int]func(interface{}))
	p.waitingResponse = make(map[int]*Response)
	p.client.Close()
}
