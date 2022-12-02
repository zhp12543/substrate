package primitives

import (
	"github.com/zhp12543/substrate/wsrpc/codec"
	"github.com/zhp12543/substrate/wsrpc/utils"
	"encoding/hex"

)

type Hash struct {
	value     []byte
	bitLength int
}

func (h *Hash) BitLength() int {
	return h.bitLength
}

func (h *Hash) EncodedLength() int {
	return h.bitLength / 8
}

func (h *Hash) String() string {
	return h.Hex()
}

func (h *Hash) Hex() string {
	d := hex.EncodeToString(h.Bytes())
	return utils.HexAddPrefix(d)
}

func (h *Hash) Bytes() []byte {
	return h.value
}

func (h *Hash) Equal(c codec.Codec) bool {
	return h.String() == c.String()
}

type Hash160 struct {
	Hash
}
type Hash256 struct {
	Hash
}
type Hash512 struct {
	Hash
}

func NewHash(data string) *Hash {
	h := Hash{}
	h.bitLength = 256
	b, err := utils.HexToBytes(data)
	if err != nil {
		h.value = []byte{}
	} else {
		h.value = b[:h.bitLength/8]
	}
	return &h
}

func NewHash160(data string) *Hash160 {
	h := Hash160{}
	h.bitLength = 160
	b, err := utils.HexToBytes(data)
	if err != nil {
		h.value = []byte{}
	} else {
		h.value = b[:h.bitLength/8]
	}
	return &h
}

func NewHash256(data string) *Hash256 {
	h := Hash256{}
	h.bitLength = 256
	b, err := utils.HexToBytes(data)
	if err != nil {
		h.value = []byte{}
	} else {
		h.value = b[:h.bitLength/8]
	}
	return &h
}

func NewHash512(data string) *Hash512 {
	h := Hash512{}
	h.bitLength = 512
	b, err := utils.HexToBytes(data)
	if err != nil {
		h.value = []byte{}
	} else {
		h.value = b[:h.bitLength/8]
	}
	return &h
}
