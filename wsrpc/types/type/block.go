package _type

type Block struct {
	Header     *BlockHeader  `json:"header"`
	Extrinsics []interface{} `json:"extrinsics"`
	Events            string `json:"events"`
}

type BlockHeader struct {
	ParentHash     string `json:"parentHash"`
	BlockHash      string `json:"blockHash"`
	Number         string `json:"number"`
	StateRoot      string `json:"stateRoot"`
	ExtrinsicsRoot string `json:"extrinsicsRoot"`
	Digest         Digest `json:"digest"`
	BlockTime      int64  `json:"blockTime"`
}

type SignedBlock struct {
	Block         *Block      `json:"block"`
	Justification interface{} `json:"justification"`
}

// 解析extrin
type TransferInfo struct {
	CallCode string
	Amount   string
	To       string
}

type ParamValues struct {
	ValueItem []*Calls `json:"value"`
}

type Calls struct {
	CallArgs     []*CallArg `json:"params"`
	CallFunction string     `json:"call_name"`
	CallIndex    string     `json:"call_index"`
	CallModule   string     `json:"call_module"`
}

type CallArg struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"value_raw"`
}

// 每个Extrinsics的交易信息
type ExtrinsicsInfo struct {
	TxId     string         `json:"txId"`
	FromAddr string         `json:"fromAddr"`
	Fee      string         `json:"fee"`
	IsBatch  bool           `json:"isBatch"`
	Transfer []TransferInfo `json:"transfer"`
	Status   bool           `json:"status"`
}
