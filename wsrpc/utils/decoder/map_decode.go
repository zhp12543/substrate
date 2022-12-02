package decoder

import (
	"github.com/zhp12543/substrate/wsrpc/types/primitives"
	"github.com/zhp12543/substrate/wsrpc/utils"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"regexp"
	"strconv"
)

func MapDecode(input interface{}, out interface{}) error {
	config := mapstructure.DecoderConfig{
		Result: out,
	}

	config.DecodeHook = mapstructure.ComposeDecodeHookFunc(handleHash, handleUint)
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	err = decoder.Decode(input)
	return nil
}

func handleHash(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	re := regexp.MustCompile(`^Hash(\d*)$`)
	matched := re.FindStringSubmatch(t.Name())
	if f.Kind() == reflect.String && t.Kind() == reflect.Struct && len(matched) > 0 {
		str := data.(string)
		var hash interface{}
		switch matched[1] {
		case "160":
			hash = primitives.NewHash160(str)
		case "256":
			hash = primitives.NewHash256(str)
		case "512":
			hash = primitives.NewHash512(str)
		case "":
			hash = primitives.NewHash256(str)
		}
		return hash, nil
	}
	return data, nil
}

func handleUint(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String && t.Kind() == reflect.Uint64 {
		str := data.(string)
		if utils.IsHex(str) {
			d := utils.HexStripPrefix(str)
			return strconv.ParseUint(d, 16, 64)
		} else {
			return strconv.Atoi(str)
		}
	}
	return data, nil
}
