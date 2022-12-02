package test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/zhp12543/substrate/config"
	v11 "github.com/zhp12543/substrate/model/v11"
	"github.com/zhp12543/substrate/rpc"
	"github.com/zhp12543/substrate/state"
	codes "github.com/itering/scale.go"
	"github.com/itering/scale.go/source"
	"github.com/itering/scale.go/types"
	"github.com/itering/scale.go/utiles"
	"testing"
)

var (
	client *rpc.Client
)

func init()  {
	var err error
	client, err = rpc.New("http://127.0.0.1:9993", "", "")
	if err != nil{
		panic(err)
	}
}

func Test_ParseExtrinsic(t *testing.T) {
	extrinsic := "0x390284cfc3f4cc12c542dc4e68f301da3477955cab60f3576847ba78835dc6405d474700bc52397521de4e78eb6a1569739cc55ddcd7be4fcc735f187d1123623ccb38cd46e1b429483f240e87f9fd41e10c02df7df5a70c77e330f3f54d20ffd374930c00950d0005005e1a84bdd0b5e1f337b60d238f56c65a43969688e3ea8ab9bcfe1d63a1ec993907801e1d8fb9"
	e := codes.ExtrinsicDecoder{}
	option := types.ScaleDecoderOption{Metadata: &client.Metadata.Metadata}
	e.Init(types.ScaleBytes{Data: utiles.HexToBytes(extrinsic)}, &option)
	e.Process()
	bb, err := json.Marshal(e.Value)
	if err != nil {
		panic(err)
	}
	var resp v11.ExtrinsicDecodeResponse
	err = json.Unmarshal(bb, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bb))
}
func Test_parseEvent(t *testing.T) {
	var (
		err  error
		key  string
		resp []byte
	)
	blockHash := "0xcbdb8536a723abfedad1faa70e845da65b579260347e2681b64f7eff8619a0fe"
	key, err = state.CreateStorageKey(client.Metadata, "System", "Events", nil, nil)
	if err != nil {
		panic(err)
	}
	resp, err = client.Rpc.SendRequest("state_getStorageAt", []interface{}{key, blockHash})
	if err != nil || len(resp) <= 0 {
		panic(err)
	}
	eventsHex := string(resp)
	//解析events
	option := types.ScaleDecoderOption{Metadata: &client.Metadata.Metadata, Spec: client.SpecVersion}
	ccHex := config.CoinEventType[client.CoinType]
	cc, _ := hex.DecodeString(ccHex)
	types.RegCustomTypes(source.LoadTypeRegistry(cc))
	e := codes.EventsDecoder{}
	e.Init(types.ScaleBytes{Data: utiles.HexToBytes(eventsHex)}, &option)
	e.Process()
	data, err1 := json.Marshal(e.Value)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(string(data))
}
