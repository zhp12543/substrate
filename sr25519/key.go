package sr25519

import (
	"encoding/hex"
	"errors"
	"fmt"
	sr25519 "github.com/ChainSafe/go-schnorrkel"
	"github.com/zhp12543/substrate/ss58"
)


type KeyPair struct {
	Wif string
	Address string
}

func GenerateKey()([]byte,[]byte, error){
	secret,err:=sr25519.GenerateMiniSecretKey()
	if err != nil {
		return nil, nil, err
	}
	if len(secret.Encode())!=32  {
		return nil, nil, errors.New("private key or public key length i not equal 32")
	}
	priv:=secret.Encode()
	pub:=secret.Public().Encode()
	return priv[:],pub[:],nil
}

func CreateAddress(pubKey,prefix []byte)(string,error){
	return ss58.Encode(pubKey,prefix)
}

func PrivateKeyToAddress(privateKey,prefix []byte)(string,error){
	var p [32]byte
	copy(p[:],privateKey[:])
	secret,err:=sr25519.NewMiniSecretKeyFromRaw(p)
	if err != nil {
		panic(err)
	}

	pub:=secret.Public().Encode()
	return ss58.Encode(pub[:],prefix)
}

func PrivateKeyToHex(privateKey []byte)(string,error){
	if len(privateKey)!=32 {
		return "",errors.New("private key length is not equal 32")
	}
	privHex:=hex.EncodeToString(privateKey)
	return "0x"+privHex,nil
}
// todo
func PrivateKeyToWif(privateKey []byte)(string,error){
	if len(privateKey)!=32 {
		return "",errors.New("private key length is not equal 32")
	}
	return "",nil
}

func Sign(privateKey, message []byte)([]byte, error){
	var sigBytes []byte
	var key, nonce [32]byte
	copy(key[:], privateKey[:32])
	signContext := sr25519.NewSigningContext([]byte("substrate"),message)
	fmt.Println(len(privateKey))
	if	len(privateKey) == 32 {	// Is seed
		sk,err := sr25519.NewMiniSecretKeyFromRaw(key)
		if err != nil {
			return nil, err
		}
		expandSk := sk.ExpandEd25519()
		sig, err := expandSk.Sign(signContext)
		if err != nil {
			return nil, err
		}

		varifySigContent := sr25519.NewSigningContext([]byte("substrate"),message)
		if ok := sk.Public().Verify(sig, varifySigContent); !ok {
			return nil, errors.New("verify sign error")
		}
		sbs := sig.Encode()
		sigBytes = sbs[:]
	}else if len(privateKey) == 64 {		//Is private key
		copy(nonce[:],privateKey[32:])
		sk := sr25519.NewSecretKey(key, nonce)
		sig,err := sk.Sign(signContext)
		if err != nil {
			return nil, fmt.Errorf("sr25519 sign error,err=%v",err)
		}

		pub, _:= sk.Public()
		if ok := pub.Verify(sig, sr25519.NewSigningContext([]byte("substrate"), message)); !ok{
			return nil, errors.New("verify sign error")
		}
		sbs := sig.Encode()
		sigBytes = sbs[:]
	}
	return sigBytes[:],nil
}
