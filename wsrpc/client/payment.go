package client

import (
	"github.com/zhp12543/substrate/wsrpc/jsonrpc"
	"github.com/zhp12543/substrate/wsrpc/types/primitives"
	"errors"
)

func createPayment(p *jsonrpc.WsProvider) *payment {
	pa := payment{}
	pa.provider = p
	pa.section = "payment"
	return &pa
}

type payment struct {
	rpcBase
}

func (p *payment) TxFee(hash *primitives.Hash256) (float64, error) {
	var params []interface{} = nil
	if hash != nil {
		params = []interface{}{hash.String()}
	}
	result, err := p.call("queryInfo", params)
	if err != nil {
		return 0.0, err
	}
	if result == nil {
		return 0.0, errors.New("get tx fee is nil")
	}

	return result.(float64), err
}
