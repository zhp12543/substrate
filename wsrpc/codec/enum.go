package codec

import (
	"reflect"
)

type Enum uint8

type EnumType struct {
	Index uint8
	Def   []interface{}
}

func (e *EnumType) Type() interface{} {
	return e.Def[e.Index]
}

func EncodeEnumType(e interface{}) ([]byte, error) {
	et := e.(EnumType)
	sub, err := Encode(et.Type())
	if err != nil {
		return nil, err
	}
	return append([]byte{uint8(et.Index)}, sub...), nil
}

func DecodeEnumType(b []byte, target reflect.Value) (*ByteInfo, error) {
	fIdx := target.FieldByName("Index")
	fIdx.SetUint(uint64(b[0]))

	fDef := target.FieldByName("Def")
	t := fDef.Index(int(b[0]))

	_, err := Decode(b[1:], &t)
	return nil, err
}
