package codec

import (
	"reflect"
)

type Option struct {
	Value interface{}
}

func (o *Option) Empty() bool {
	return o.Value == nil
}

func EncodeOption(data interface{}) ([]byte, error) {
	d := data.(Option)
	if d.Empty() {
		return []byte{0}, nil
	}
	sub, err := Encode(d.Value)
	if err != nil {
		return nil, err
	}
	return append([]byte{1}, sub...), nil
}

func DecodeOption(b []byte, target reflect.Value) (*ByteInfo, error) {
	if b[0] == 0 {
		v := reflect.ValueOf(Option{Value: nil})
		target.Set(v)
		return &ByteInfo{Offset: 0, Len: 1}, nil
	}

	v := target.FieldByName("Value")
	info, err := Decode(b[1:], &v)
	if err != nil {
		return nil, err
	}
	return &ByteInfo{Offset: 0, Len: info.End() + 1}, nil
}
