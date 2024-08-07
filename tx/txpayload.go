package tx

import (
	"encoding/hex"
)

type TxPayLoad struct {
	Method []byte
	Era []byte
	Nonce []byte
	Fee []byte
	SpecVersion []byte
	TransactionVersion []byte
	GenesisHash []byte
	BlockHash []byte
}

func (t TxPayLoad) ToBytesString (methodPrefix []byte) string {
	payload := make([]byte, 0)
	//payload = append(payload,[]byte{0x04,0x00}...)
	payload = append(payload, t.Method...)
	payload = append(payload, t.Era...)
	payload = append(payload, t.Nonce...)
	payload = append(payload, t.Fee...)
	if len(methodPrefix) > 0{
		// ksm 添加00
		payload = append(payload, methodPrefix...)
	}
	payload = append(payload, t.SpecVersion...)
	payload = append(payload,t.TransactionVersion...)
	payload = append(payload, t.GenesisHash...)
	payload = append(payload, t.BlockHash...)
	if len(methodPrefix) > 0{
		// ksm 添加00
		payload = append(payload, methodPrefix...)
	}

	return hex.EncodeToString(payload)
}