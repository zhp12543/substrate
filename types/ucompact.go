package types

import (
	"github.com/zhp12543/substrate/scale"
	"math/big"
)

type UCompact big.Int

func NewUCompact(value *big.Int) UCompact {
	return UCompact(*value)
}

func (u *UCompact) Int64() int64 {
	i := big.Int(*u)
	return i.Int64()
}

func NewUCompactFromUInt(value uint64) UCompact {
	return NewUCompact(new(big.Int).SetUint64(value))
}

func (u *UCompact) Decode(decoder scale.Decoder) error {
	ui, err := decoder.DecodeUintCompact()
	if err != nil {
		return err
	}

	*u = UCompact(*ui)
	return nil
}

func (u UCompact) Encode(encoder scale.Encoder) error {
	err := encoder.EncodeUintCompact(big.Int(u))
	if err != nil {
		return err
	}
	return nil
}