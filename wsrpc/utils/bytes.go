package utils

import (
	"encoding/hex"
)

func BytePad(b []byte, w int, suffix bool) []byte {
	l := len(b)
	if l >= w {
		return b
	}
	bb := make([]byte, w)
	if suffix {
		for i := 0; i < l; i++ {
			bb[i] = b[i]
		}
	} else {
		for i := w - 1; i >= w-l; i-- {
			bb[i] = b[l-w+i]
		}
	}
	return bb
}

func ByteToHex(data []byte) string {
	s := hex.EncodeToString(data)
	return HexAddPrefix(s)
}
