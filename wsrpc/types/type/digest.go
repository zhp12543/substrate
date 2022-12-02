package _type

type Digest struct {
	Logs []string `json:"logs"`
}

type DigestItem struct {
}

type Signature []byte
type Seal map[uint64]Signature
