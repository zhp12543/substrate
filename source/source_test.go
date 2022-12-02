package source_test

import (
	"encoding/hex"
	"fmt"
	"github.com/zhp12543/substrate/config"
	"io/ioutil"
	"testing"
)

func TestLoadTypeRegistry(t *testing.T) {
	cc, err := ioutil.ReadFile(fmt.Sprintf("%s.json", config.Polkadot))
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(cc))
}
