package tx

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/zhp12543/substrate/scale"
	"github.com/zhp12543/substrate/types"
	"math/big"
)

const (
	SigningBitV4 = byte(0x84)
)

type MethodTransfer struct {
	DestPubkey []byte
	Amount     []byte
}

func NewMethodTransfer(pubkey string, amount *big.Int) (*MethodTransfer, error) {
	pubBytes, err := hex.DecodeString(pubkey)
	if err != nil || len(pubBytes) != 32 {
		return nil, errors.New("invalid dest public key")
	}

	if amount.Cmp(big.NewInt(0)) <= 0 {
		return nil, errors.New("amount <= 0")
	}
	uAmount := types.UCompact(*new(big.Int).Set(amount))
	var buffer = bytes.Buffer{}

	s := scale.NewEncoder(&buffer)
	errA := uAmount.Encode(*s)
	if errA != nil {
		return nil, fmt.Errorf("encode amount error,Err=[%v]", errA)
	}

	return &MethodTransfer{
		DestPubkey: pubBytes,
		Amount:     buffer.Bytes(),
	}, nil
}

func (mt MethodTransfer) ToBytes(callId string, accPrefix, accSuffix []byte) ([]byte, error) {

	if mt.DestPubkey == nil || len(mt.DestPubkey) != 32 || mt.Amount == nil || len(mt.Amount) == 0 {
		return nil, errors.New("invalid method")
	}

	ret, _ := hex.DecodeString(callId)
	if  len(accSuffix) > 0{
		ret = append(ret, accSuffix...)
	}

	if accPrefix != nil {
		ret = append(ret, accPrefix...)
	}

	ret = append(ret, mt.DestPubkey...)
	ret = append(ret, mt.Amount...)

	return ret, nil
}