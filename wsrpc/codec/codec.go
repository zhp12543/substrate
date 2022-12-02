package codec

type Codec interface {
	EncodedLength() int
	Hex() string
	Bytes() []byte
	String() string
	Equal(b interface{}) bool
	Empty() bool
}
