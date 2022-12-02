package client

import (
	"github.com/zhp12543/substrate/wsrpc/codec"
	"github.com/zhp12543/substrate/wsrpc/jsonrpc"
	_type "github.com/zhp12543/substrate/wsrpc/types/type"
	"github.com/zhp12543/substrate/wsrpc/utils"
)

func createAuthor(p *jsonrpc.WsProvider) *author {
	a := author{}
	a.provider = p
	a.section = "author"
	return &a
}

type author struct {
	rpcBase
}

func (a *author) PendingExtrinsics() ([]_type.Extrinsic, error) {
	result, err := a.call("pendingExtrinsics", nil)
	if err != nil {
		return nil, err
	}
	var extrinsics []_type.Extrinsic
	for _, item := range result.([]string) {
		b, _ := utils.HexToBytes(item)
		ex := _type.Extrinsic{}
		_, err := codec.Decode(b, &ex)
		if err != nil {
			return nil, err
		}
		extrinsics = append(extrinsics, ex)
	}
	return extrinsics, nil
}

func (a *author) SubmitAndWatchExtrinsic(signture string) {
	a.call("submitAndWatchExtrinsic", nil)
}

func (a *author) SubmitExtrinsic(ex _type.Extrinsic) {
	// todo
}
