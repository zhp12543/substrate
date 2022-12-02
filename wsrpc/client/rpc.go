package client

import "github.com/zhp12543/substrate/wsrpc/jsonrpc"

type rpc struct {
	System  *system
	State   *state
	Chain   *chain
	Author  *author
	Payment *payment
}

type rpcBase struct {
	provider *jsonrpc.WsProvider
	section  string
}

func (r *rpcBase) call(method string, params []interface{}) (interface{}, error) {
	resp, err := r.provider.Call(r.section+"_"+method, params)
	return resp.Result, err
}

func (r *rpcBase) subscribe(method string, params []interface{}, callback func(interface{})) (int, error) {
	return r.provider.Subscribe(r.section+"_"+method, params, callback)
}
