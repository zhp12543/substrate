package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOption(t *testing.T) {
	original := Option{
		Value: "hello",
	}
	b, _ := Encode(original)

	target := Option{Value: ""}
	info, _ := Decode(b, &target)
	assert.Equal(t, original.Value, target.Value)
	assert.Equal(t, uint64(0), info.Offset)
	assert.Equal(t, uint64(len(b)), info.Len)
}
