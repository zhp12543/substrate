package codec

import (
	"github.com/zhp12543/substrate/wsrpc/utils"
	"encoding/binary"
	"math"
)

func AddLengthPrefix(data []byte) []byte {
	l := len(data)
	var prefix []byte
	if l <= math.MaxUint8 {
		prefix = []byte{uint8(l << 2)}
	} else if l <= math.MaxUint16 {
		prefix = utils.UintToFixedByte(uint64(l<<2+0x01), 2)
	} else if l <= math.MaxUint32 {
		prefix = utils.UintToFixedByte(uint64(l<<2+0x02), 4)
	} else {
		prefix := utils.UintToByte(uint64(l))
		bSize := len(prefix)
		for i := len(prefix); i > 0; i++ {
			if prefix[i] == 0 {
				bSize = i
			} else {
				break
			}
		}
		prefix = []byte{uint8((bSize-4)<<2 + 2), uint8(bSize)}
	}
	return append(prefix, data...)
}

type ByteInfo struct {
	Offset uint64
	Len    uint64
}

func (b *ByteInfo) End() uint64 {
	return b.Offset + b.Len
}

func GetBytesInfo(data []byte) *ByteInfo {
	var offset, l uint64

	flag := data[0] & 0x03
	switch flag {
	case 0x00:
		offset = 1
		l = uint64(data[0] >> 2)
	case 0x01:
		d := utils.BytePad(data[:2], 8, true)
		offset = 2
		l = binary.LittleEndian.Uint64(d) >> 2
	case 0x02:
		d := utils.BytePad(data[:4], 8, true)
		offset = 4
		l = binary.LittleEndian.Uint64(d) >> 2
	default:
		offset = uint64(data[0]>>2) + 4 + 1
		d := utils.BytePad(data[1:offset], 8, true)
		l = binary.LittleEndian.Uint64(d)
	}

	return &ByteInfo{
		Offset: offset,
		Len:    l,
	}
}
