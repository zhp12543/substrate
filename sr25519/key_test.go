package sr25519

import (
	"encoding/hex"
	"fmt"
	sr25519 "github.com/ChainSafe/go-schnorrkel"
	"github.com/zhp12543/substrate/config"
	"github.com/zhp12543/substrate/ss58"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/blake2b"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	secret := ""
	priv,err:=hex.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	//fmt.Println(len(priv))
	if len(priv)!=32 {
		return
	}
	var s [32]byte
	copy(s[:],priv)
	fmt.Println(s)
	key,err:=sr25519.NewMiniSecretKeyFromRaw(s)
	if err != nil {
		panic(err)
	}
	p:=key.ExpandEd25519().Encode()
	fmt.Println("=============")
	fmt.Println(hex.EncodeToString(p[:]))
	_,pubK := key.ExpandEd25519(), key.Public()
	pp:=pubK.Encode()
	fmt.Println(hex.EncodeToString(pp[:]))
	pub:=pubK.Encode()
	fmt.Println(ss58.Encode(pub[:],config.PolkadotPrefix))

}
var(
	ssPrefix = []byte{0x53, 0x53, 0x35, 0x38, 0x50, 0x52, 0x45}
)
func TestCreateAddress2(t *testing.T) {
	address:="Dj1nP5Ebtjo5GZhksnh1zyDWgAcmiuqmXC6PEe6YzD7JNT6"
	data:=base58.Decode(address)
	fmt.Println(data)
	fmt.Println(data[:33])
	fmt.Println(hex.EncodeToString(data[1:33]))
	var d []byte
	d = append(d,ssPrefix...)
	d = append(d,data[:33]...)
	s:=blake2b.Sum512(d)
	fmt.Println(s)
}

func TestCreateAddress3(t *testing.T) {
	priv,pub,err:=GenerateKey()
	fmt.Println(len(priv),len(pub))
	if err != nil {
		fmt.Println(11)
		panic(err)
	}
	private,err:=PrivateKeyToHex(priv)
	if err != nil {
		panic(err)
	}
	address,err:=CreateAddress(pub,config.SubstratePrefix)
	if err != nil {
		panic(err)
	}
	fmt.Println(private)
	fmt.Println(address)
	/*
	F2WdpYJw37tjBWLAgUZQRN1hxhwUF285L1L3NaqhWKw3tUU
	*/
}

func TestSign(t *testing.T) {
	secret:=""
	priv,err:=hex.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	//fmt.Println(len(priv))
	if len(priv)!=32 {
		return
	}
	var s [32]byte
	copy(s[:],priv)
	fmt.Println(s)
	key,err:=sr25519.NewMiniSecretKeyFromRaw(s)
	if err != nil {
		panic(err)
	}

	pub,_:=key.ExpandEd25519().Public()
	p:=pub.Encode()
	fmt.Println(hex.EncodeToString(p[:]))
	//priv,err:=hex.DecodeString(secret)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Println(len(priv))
	//if len(priv)!=32 {
	//	return
	//}
	//
	//message:=[]byte("Test123")
	//sign,err:=Sign(priv,message)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(hex.EncodeToString(sign))

}
func TestGenerateKey(t *testing.T) {

}

func TestPrivateKeyToHex(t *testing.T) {
	privateKey,_:=hex.DecodeString("")
	sig,_:=hex.DecodeString("f67567cbd08a43da678c9534d607034a5b5fc8c7b89104b8e7da448a3294375e542ce2e595bfa2cce3a4c180d3a85aaab297dadd36b60318b7ffe9f34b660d8b")
	var key,nonce [32]byte
	copy(key[:],privateKey[:32])
	copy(nonce[:],privateKey[32:])
	fmt.Println(len(privateKey))
	fmt.Println(key)
	fmt.Println(nonce)
	secret:=sr25519.NewSecretKey(key,nonce)
	var sigBytes [64]byte
	copy(sigBytes[:],sig)
	fmt.Println(len(sig))
	signs:=sr25519.Signature{}
	signs.Decode(sigBytes)
	pub,err:=secret.Public()
	if err != nil {
		fmt.Println(err)
		return
	}
	isOK:=pub.Verify(&signs,sr25519.NewSigningContext([]byte("substrate"),[]byte("Test123")))
	fmt.Println(isOK)
}