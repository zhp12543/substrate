package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type d struct {
	Name     string
	Age      int
	Approved bool
}

func TestEnumType(t *testing.T) {
	original := EnumType{
		Index: 1,
		Def: []interface{}{
			nil,
			d{
				Name:     "hello",
				Age:      33,
				Approved: true,
			},
		},
	}
	b, _ := Encode(original)

	target := EnumType{
		Def: []interface{}{
			nil,
			d{},
		},
	}
	Decode(b, &target)
	assert.Equal(t, original.Index, target.Index)
	assert.Equal(t, original.Type(), target.Type())
}
