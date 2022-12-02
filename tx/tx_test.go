package tx

import (
	"github.com/zhp12543/substrate/config"
	"github.com/zhp12543/substrate/rpc"
	"github.com/zhp12543/substrate/ss58"
	"encoding/hex"
	"fmt"
	sr255191 "github.com/ChainSafe/go-schnorrkel"
	"math/big"
	"testing"
)

var (
	client *rpc.Client
)

func init()  {
	var err error
	client, err = rpc.New("http://127.0.0.1:9993", "", "")
	if err != nil{
		panic(err)
	}
}

func NewKeypairFromSeed(seed string) ([]byte, []byte, error) {
	s, e := hex.DecodeString(seed)
	if e != nil{
		return nil, nil, e
	}

	buf := [32]byte{}
	copy(buf[:], s)
	msc, err := sr255191.NewMiniSecretKeyFromRaw(buf)
	if err != nil {
		return nil, nil, err
	}

	priv := msc.Encode()
	pub, err2  := msc.ExpandEd25519().Public()
	if err2 != nil {
		return nil, nil, err2
	}

	pp := pub.Encode()

	return priv[:], pp[:], nil

}

func TestNewMethodTransfer(t *testing.T) {
	seed := ""
	fmt.Println(seed)
	//fromPriv, fromPub, err := sr25519.GenerateKey()
	fromPriv, fromPub, err := NewKeypairFromSeed(seed)
	fmt.Println("from:", hex.EncodeToString(fromPriv), hex.EncodeToString(fromPub), err)

	//toPriv, toPub, err1 := sr25519.GenerateKey()
	//fmt.Println("to:", hex.EncodeToString(toPriv), hex.EncodeToString(toPub), err1)

	var (
		from = ""
		to = "1338f2WYvtbydpruikkCyZ7DhnMdJ742m72Vom5x9Dw8b6Dw"
	)

	from, err = ss58.PublicKeyToAddress(fromPub)
	fmt.Println("from addr:", from, err)

	//to, err = PublicKeyToAddress(toPub)
	//fmt.Println("to addr:", to, err)

	fmt.Println("from public eq:", hex.EncodeToString(fromPub) == ss58.AddressToPublicKey(from))
	//fmt.Println("to public eq:", hex.EncodeToString(toPub) == AddressToPublicKey(to))


	//client.GetFinalizedHead()
	//block, err0 := client.GetBlockByNumber(-1)
	//fmt.Println("get last block:", block, err0)

	//lastBlockNum := block.Height
	//blockHash := block.BlockHash
	specVersion, txVersion := uint32( client.SpecVersion ), uint32( client.TransactionVersion )
	genesisHash, _ := client.GetGenesisHash()
	var nonce uint64 = 1
	lastBlockNum := 567
	blockHash := genesisHash
	var amount int64 = (1-0.0153) * 10000000000
	fmt.Println("trans amount:", amount/10000000000)

	trans := CreateTransaction(from, to, new(big.Int).SetInt64(amount), nonce, 0)
	trans.SetGenesisHashAndBlockHash(genesisHash, blockHash, uint64(lastBlockNum))
	trans.SetSpecVersionAndCallId(specVersion, txVersion, config.CallIdDot)

	_, msg, err2 := trans.CreateEmptyTransactionAndMessage()
	fmt.Println("create tx:", msg, err2)

	sige, err3 := trans.SignTransaction(hex.EncodeToString(fromPriv), msg)
	fmt.Println("sign:", sige, err3)

	finalTx, err4 := trans.GetSignTransaction(sige, 1)

	fmt.Println("final tx:", finalTx, err4)
	fmt.Println("txid unsubmit:", rpc.CreateTxHash(finalTx))

	txidBytes,err:=client.Rpc.SendRequest("author_submitExtrinsic",[]interface{}{finalTx})
	if err != nil {
		panic(err)
	}
	txid := string(txidBytes)
	fmt.Println(txid)
}