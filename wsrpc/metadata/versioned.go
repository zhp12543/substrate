package metadata

import "github.com/zhp12543/substrate/wsrpc/codec"

type Versioned struct {
	MagicNumber uint32
	Metadata    codec.EnumType
}

func (v *Versioned) Version() uint8 {
	return v.Metadata.Index
}

type PlainType struct {
	Value string
}

type MapType struct {
	Key      string
	Value    string
	IsLinked bool
}

type DoubleMapType struct {
	Key1      string
	Key2      string
	Value     string
	KeyHasher string
}

func CreateMetadata() Versioned {
	return Versioned{
		Metadata: codec.EnumType{
			Def: []interface{}{},
		},
	}
}
