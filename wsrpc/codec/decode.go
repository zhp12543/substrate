package codec

import (
	"encoding/binary"
	"errors"
	"reflect"
)

func Decode(data []byte, s interface{}) (*ByteInfo, error) {
	val := reflect.Indirect(reflect.ValueOf(s))
	// if out source is a reflect.Value reference, use it directly
	if val.Type().String() == "reflect.Value" {
		val = *(s.(*reflect.Value))
	}

	switch val.Kind() {
	case reflect.String:
		return decodeString(data, val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return decodeInt(data, val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return decodeUint(data, val)
	case reflect.Bool:
		return decodeBool(data, val)
	case reflect.Slice:
		return decodeSlice(data, val)
	case reflect.Struct:
		return decodeStruct(data, val)
	case reflect.Interface:
		return decodeInterface(data, val)
	case reflect.Ptr:
		val = val.Elem()
		return Decode(data, &val)
	}
	return nil, errors.New("can not decode with type " + val.Kind().String())
}

func decodeString(b []byte, val reflect.Value) (*ByteInfo, error) {
	info := GetBytesInfo(b)
	v := string(b[info.Offset:info.End()])
	val.SetString(v)
	return info, nil
}

func decodeUint(b []byte, val reflect.Value) (*ByteInfo, error) {
	info := &ByteInfo{
		Offset: 0,
		Len:    uint64(val.Type().Size()),
	}
	b = b[:info.Len]
	n, _ := binary.Uvarint(b)
	val.SetUint(n)
	return info, nil
}

func decodeInt(b []byte, val reflect.Value) (*ByteInfo, error) {
	info := &ByteInfo{
		Offset: 0,
		Len:    uint64(val.Type().Size()),
	}
	b = b[:info.Len]
	n, _ := binary.Varint(b)
	val.SetInt(n)
	return info, nil
}

func decodeBool(b []byte, val reflect.Value) (*ByteInfo, error) {
	info := &ByteInfo{
		Offset: 0,
		Len:    1,
	}
	if b[0] == 1 {
		val.SetBool(true)
	} else {
		val.SetBool(false)
	}
	return info, nil
}

func decodeSlice(b []byte, val reflect.Value) (*ByteInfo, error) {
	info := GetBytesInfo(b)
	b = b[info.Offset:]
	l := int(info.Len)
	valSlice := reflect.MakeSlice(val.Type(), l, l)
	for i := 0; i < l; i++ {
		v := valSlice.Index(i)
		subInfo, err := Decode(b, &v)
		if err != nil {
			return nil, err
		}
		b = b[subInfo.End():]
	}
	val.Set(valSlice)
	return info, nil
}

func decodeStruct(b []byte, val reflect.Value) (*ByteInfo, error) {
	t := val.Type()
	name := t.Name()
	if decoder, ok := decodeMap[name]; ok {
		return decoder(b, val)
	}

	info := GetBytesInfo(b)
	b = b[info.Offset:]
	for i := 0; i < val.Type().NumField(); i++ {
		f := val.Field(i)
		subInfo, err := Decode(b, &f)
		if err != nil {
			return nil, err
		}
		b = b[subInfo.End():]
	}
	return info, nil
}

func decodeInterface(b []byte, val reflect.Value) (*ByteInfo, error) {
	newVal := reflect.New(val.Elem().Type()).Elem()
	offset, err := Decode(b, &newVal)
	if err != nil {
		return nil, err
	}
	val.Set(newVal)
	return offset, nil
}
