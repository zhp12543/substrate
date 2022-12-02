package codec

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func makeRandomByte(l int) []byte {
	var b = make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = uint8(rand.Intn(256))
	}
	return b
}

func TestAddLengthPrefix(t *testing.T) {
	b16 := makeRandomByte(300)
	data := []struct {
		input  []byte
		expect []byte
	}{
		{[]byte{12, 13}, []byte{2 << 2, 12, 13}},
		{b16, append([]byte{0xb1, 0x4}, b16...)},
	}

	for _, d := range data {
		b := AddLengthPrefix(d.input)
		assert.Equal(t, d.expect, b)
	}
}

func TestGetBytesInfo(t *testing.T) {
	data := []struct {
		input  []byte
		expect []uint64
	}{
		{[]byte{252}, []uint64{1, 63}},
		{[]byte{253, 7}, []uint64{2, 511}},
		{[]byte{254, 255, 3, 0}, []uint64{4, 0xffff}},
		{[]byte{3, 249, 255, 255, 255}, []uint64{5, 0xfffffff9}},
	}

	for _, d := range data {
		info := GetBytesInfo(d.input)
		assert.Equal(t, d.expect, []uint64{info.Offset, info.Len})
	}
}
