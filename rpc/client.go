package rpc

import (
	"github.com/zhp12543/substrate/config"
	"github.com/zhp12543/substrate/util"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	codes "github.com/itering/scale.go"
	"github.com/itering/scale.go/source"
	"github.com/itering/scale.go/types"
	"golang.org/x/crypto/blake2b"
	"math/big"
	"net/http"
	"strings"
)

type Client struct {
	Rpc                *util.RpcClient
	Metadata           *codes.MetadataDecoder
	CoinType           string
	SpecVersion        int
	TransactionVersion int
	genesisHash        string
}

func CreateTxHash(extrinsic string) string {
	data, _ := hex.DecodeString(util.RemoveHex0x(extrinsic))
	d := blake2b.Sum256(data)
	return "0x" + hex.EncodeToString(d[:])
}

func New(url, user, password string) (*Client, error) {
	client := new(Client)
	if strings.HasPrefix(url, "wss") {
		//todo 连接websocket
		return client, errors.New("do not support websocket")
	}
	// 新建一个全局的http请求对象
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic("http请求对象创建失败，请检查yaml文件节点配置是否正确")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client.Rpc = util.New(url, user, password, req)
	//初始化运行版本
	err = client.InitRuntimeVersion()
	if err != nil {
		return nil, err
	}
	client.registerTypes()
	return client, nil
}

func (client *Client) InitMetaData() error {
	metadataBytes, err := client.Rpc.SendRequest("state_getMetadata", []interface{}{})
	if err != nil {
		return fmt.Errorf("rpc get metadata error,err=%v", err)
	}
	metadata := string(metadataBytes)
	metadata = util.RemoveHex0x(metadata)
	data, err := hex.DecodeString(metadata)
	if err != nil {
		return err
	}
	m := codes.MetadataDecoder{}
	m.Init(data)
	if err := m.Process(); err != nil {
		return fmt.Errorf("parse metadata error,err=%v", err)
	}
	client.Metadata = &m
	return nil
}

/*
注册types
*/
func (client *Client) registerTypes() {
	ccHex := config.CoinEventType[client.CoinType]
	cc, _ := hex.DecodeString(ccHex)
	types.RegCustomTypes(source.LoadTypeRegistry(cc))
}

func (client *Client) InitRuntimeVersion() error {
	data, err := client.Rpc.SendRequest("state_getRuntimeVersion", []interface{}{})
	if err != nil {
		return fmt.Errorf("init runtime version error,err=%v", err)
	}
	var result map[string]interface{}
	errJ := json.Unmarshal(data, &result)
	if errJ != nil {
		return fmt.Errorf("init runtime version error,err=%v", errJ)
	}
	client.CoinType = strings.ToLower(result["specName"].(string))
	client.TransactionVersion = int(result["transactionVersion"].(float64))
	specVersion := int(result["specVersion"].(float64))
	// metadata 会动态改变，所以通过specVersion去检测metadata的改变
	if client.SpecVersion != specVersion {
		client.SpecVersion = specVersion
		return client.InitMetaData()
	}
	client.SpecVersion = specVersion
	return nil
}

func (client *Client) GetFinalizedHead() (string, error) {
	resp, err := client.Rpc.SendRequest("chain_getFinalizedHead", []interface{}{})
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (client *Client) GetGenesisHash() (string, error) {
	if client.genesisHash != "" {
		return client.genesisHash, nil
	}

	resp, err := client.Rpc.SendRequest("chain_getBlockHash", []interface{}{0})
	if err != nil {
		return "", err
	}

	client.genesisHash = string(resp)
	return string(resp), nil
}

func (client *Client) PaymentQueryInfo(extrin string) (map[string]interface{}, error) {
	var (
		err  error
		resp []byte
	)
	resp, err = client.Rpc.SendRequest("payment_queryInfo", []interface{}{extrin})
	if err != nil || len(resp) <= 0 {
		return nil, fmt.Errorf("get payment_queryInfo error,err=%v", err)
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal payment_queryInfo error,err=%v", err)
	}
	return data, nil
}

func (client *Client) RemoveExtrinsic(hashOrExtrinsic string) error {
	var (
		respData []byte
		err      error
	)
	params := make([]interface{}, 0)
	if hashOrExtrinsic != "" {
		params = append(params, []interface{}{hashOrExtrinsic})
	}
	respData, err = client.Rpc.SendRequest("author_removeExtrinsic", params)
	if err != nil || len(respData) == 0 {
		return fmt.Errorf("remove extrinsic error,err=%v", err)
	}
	return nil
}

/**
 * SendRawTransaction sends tx to node
 * @param signture string
 * @return error
 */
func (client *Client) AuthorSubmitExtrinsic(signture string) (string, error) {
	txIdBytes, err := client.Rpc.SendRequest("author_submitExtrinsic", []interface{}{signture})
	return string(txIdBytes), err
}

func (client *Client) SystemAccountNextIndex(address string) (string, error) {
	data, err := client.Rpc.SendRequest("system_accountNextIndex", []interface{}{address})
	return string(data), err
}

func (client *Client) ChainGetBlock(hash string) (string, error) {
	params := make([]interface{}, 1)
	if hash != "" {
		params[0] = hash
	}
	data, err := client.Rpc.SendRequest("chain_getBlock", params)
	return string(data), err
}

func (client *Client) ChainGetBlockHash(height *big.Int) (string, error) {
	params := make([]interface{}, 0)
	if height != nil{
		params = append(params, height.Int64())
	}
	data, err := client.Rpc.SendRequest("chain_getBlockHash", params)
	return string(data), err
}

func (client *Client) ChainGetFinalizedHead() (string, error) {
	params := make([]interface{}, 0)
	data, err := client.Rpc.SendRequest("chain_getFinalizedHead", params)
	return string(data), err
}

func (client *Client) StateGetStorageAt(key , blockHash string) (string, error) {
	data, err := client.Rpc.SendRequest("state_getStorageAt", []interface{}{key, blockHash})
	if err != nil || len(data) <= 0 {
		return "", err
	}
	return string(data), nil
}