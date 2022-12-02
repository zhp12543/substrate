package rpc

import (
	"fmt"
	"testing"
)

var client *Client
func init() {
	var err error
	client, err = New("http://127.0.0.1:9993", "", "")
	if err != nil{
		panic(err)
	}
}

func TestCreateTxHash(t *testing.T) {
	extrinsic := "0x3902845098eb69050aae4a21b9480678d92f1bef35d906169147cd2fa3c7ea267eee3e0178dd9bb837ec5f2fe8432ba3d9cd3ece9da46f44f080d331b2fe61e6c3f9080d9e2ffe166ad0d1a3d4a7d4d9113fa4f763d6e4eb15f2f9179e49e9448fc8348200d501000500727903930e113ec1ae36807662540eadc957703a20b20b0fae0751220515817c070082357a0a"
	fmt.Println(CreateTxHash(extrinsic))
}

func TestRemoveExtrinsic(t *testing.T) {
	sign := "0x3502843e044b13db56a7fbcc2a583938cd9c1c7dc96eb116e2f762923e8ea1b145f95001261d80847bf90c3f019eb95e5250a3196318cec3567baeeda43887222d930233e80f7ee4b75b3c92852ea7813f3d82bb19e85235a7cd0997acc83035dcec3c85000c00050098962a6365b6fe9a4149e3872bd3bd0456a9d792c14f533fe7dac4e995d5dc7307808e9d9504"
	client.RemoveExtrinsic(sign)
}