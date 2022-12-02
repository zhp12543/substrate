package utils

import (
	"encoding/binary"
	"math"
)

func IntToByte(d int64) []byte {
	var b = make([]byte, binary.MaxVarintLen64)
	l := binary.PutVarint(b, d)
	return b[:l]
}

func UintToByte(d uint64) []byte {
	var size = 8
	if d <= math.MaxUint8 {
		size = 1
	} else if d <= math.MaxUint16 {
		size = 2
	} else if d <= math.MaxUint32 {
		size = 4
	}
	var b = make([]byte, binary.MaxVarintLen64)
	binary.LittleEndian.PutUint64(b, d)
	return b[:size]
}

func IntToFixedByte(d int64, size uint) []byte {
	var b = make([]byte, binary.MaxVarintLen64)
	l := binary.PutVarint(b, d)
	s := int(size)
	if l < s {
		l = s
	}
	return b[:l]
}

func UintToFixedByte(d uint64, size uint) []byte {
	var b = make([]byte, binary.MaxVarintLen64)
	binary.LittleEndian.PutUint64(b, d)
	return b[:size]
}
