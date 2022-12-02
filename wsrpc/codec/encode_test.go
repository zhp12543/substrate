package codec

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestEncode(t *testing.T) {
	data := []struct {
		input  interface{}
		expect []byte
	}{
		{"hello", []byte{0x14, 0x68, 0x65, 0x6c, 0x6c, 0x6f}},
		{true, []byte{1}},
		{false, []byte{0}},
		{uint8(math.MaxUint8), []byte{255}},
		{uint16(math.MaxUint16), []byte{255, 255}},
		{uint32(math.MaxUint32), []byte{255, 255, 255, 255}},
		{uint64(math.MaxUint64), []byte{255, 255, 255, 255, 255, 255, 255, 255}},
		{[]uint16{math.MaxUint16, math.MaxUint16}, []byte{0x10, 255, 255, 255, 255}},
		{Enum(255), []byte{255}},
		{
			struct {
				Name string
				Age  uint8
			}{
				Name: "zhex",
				Age:  38,
			},
			[]byte{0x18, 0x10, 0x7a, 0x68, 0x65, 0x78, 0x26},
		},
	}
	for _, d := range data {
		data, _ := Encode(d.input)
		assert.Equal(t, d.expect, data)
	}
}
