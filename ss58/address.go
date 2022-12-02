package ss58

import (
	"encoding/hex"
	"fmt"
)

func VerifyAddress(address string, prefix []byte) bool{
	err := VerityAddress(address, prefix)
	//fmt.Println("verify address err:", err)
	return err == nil
}

func AddressToPublicKey(address string) string {
	if address == "" {
		return ""
	}
	pub, err := DecodeToPub(address)
	if err != nil {
		fmt.Println("deccode address to pub err:", err)
		return ""
	}
	if len(pub) != 32 {
		fmt.Println("deccode address to pub err: pub len !=32")
		return ""
	}
	pubHex := hex.EncodeToString(pub)
	return pubHex
}

func PublicKeyToAddress(pub, prefix []byte) (string, error) {
	return Encode(pub, prefix)
}