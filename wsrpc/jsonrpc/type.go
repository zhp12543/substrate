package jsonrpc

type Request struct {
	ID      int         `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type Response struct {
	ID      int             `json:"id,omitempty"`
	JSONRPC string          `json:"jsonrpc"`
	Result  interface{}     `json:"result,omitempty"`
	Params  SubscribeResult `json:"params,omitempty"`
	Error   ResponseError   `json:"error,omitempty"`
}

type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SubscribeResult struct {
	Result       interface{} `json:"result"`
	Subscription int         `json:"subscription"`
}

type HttpError struct {
	Code int
	err  error
}

func (e *HttpError) Error() string {
	return e.err.Error()
}
