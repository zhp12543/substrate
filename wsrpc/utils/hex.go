package utils

import (
	"encoding/hex"
	"regexp"
	"strings"
)

const (
	HexPrefix = "0x"
)

func IsHex(data string) bool {
	if len(data) == 0 {
		return false
	}
	regex := regexp.MustCompile(`^(0x)?[\da-fA-F]*$`)
	return regex.Match([]byte(data))
}

func HexToBytes(data string) ([]byte, error) {
	data = HexStripPrefix(data)
	return hex.DecodeString(data)
}

func HexStripPrefix(data string) string {
	return strings.TrimPrefix(data, HexPrefix)
}

func HexHasPrefix(data string) bool {
	if len(data) < 2 {
		return false
	}
	return data[:2] == HexPrefix
}

func HexAddPrefix(data string) string {
	if data[:2] == HexPrefix {
		return data
	}
	return HexPrefix + data
}
