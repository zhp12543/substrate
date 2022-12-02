package polkadotrpc

import (
	"fmt"
	"github.com/zhp12543/substrate/wsrpc/client"
	"github.com/zhp12543/substrate/wsrpc/types/primitives"
	"testing"
)


var c *client.Client

func init() {
	c, _ = client.New("ws://127.1.1.0:9995/")
}

func TestCommon(t *testing.T) {
	t.Run("获取精度", testGetDecimal)
	t.Run("获取nonce", testGetNonce)
	t.Run("获取最后一个不可逆块hash", testGetFinalizedHead)
	t.Run("获取指定块高hash", testGetBlockHash)
	t.Run("获取指定块数据", testGetBlock)
	t.Run("pending 数据", testGetPending)
}

// pending 数据
func testGetPending(t *testing.T)  {
	pendingExtrinsics, err := c.RPC.Author.PendingExtrinsics()
	fmt.Println(pendingExtrinsics, err)
}


// 获取精度
func testGetDecimal(t *testing.T)  {
	chainProperties, err := c.RPC.System.Properties()
	fmt.Println(chainProperties.TokenDecimals, err)
}

// 获取nonce
func testGetNonce(t *testing.T)  {
	nonce, err := c.RPC.System.Nonce("FJaSzBUAJ1Nwa1u5TbKAFZG5MBtcUouTixdP7hAkmce2SDS")
	fmt.Println(nonce, err)
}

// 获取最后一个不可逆块hash
func testGetFinalizedHead(t *testing.T)  {
	hash, _ := c.RPC.Chain.GetFinalisedHead()
	fmt.Println(hash.String())
}


//	获取指定块高hash
func testGetBlockHash(t *testing.T)  {
	var h uint64 = 1000
	name, err := c.RPC.Chain.GetBlockHash(&h)
	fmt.Println(name, err)
}

// 获取指定块数据
func testGetBlock(t *testing.T)  {
	hash := primitives.NewHash256("0xe6cce572b0fb7e09b53d86f986ba4ac900467eb5bd07cb02b1b4721f8c2f23cd")
	blockInfo, err := c.RPC.Chain.GetBlock(hash)
	fmt.Println(*blockInfo, err)
}
