package lda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValues(t *testing.T) {
	tests := []struct {
		in  interface{}
		out interface{}
	}{
		{map[string]string{
			"k1": "v1",
			"k2": "v2",
		}, "v1"},
		{map[string]interface{}{
			"k1": "v1",
			"k2": "v2",
		}, "v1"},
	}
	for _, te := range tests {
		out := Values(te.in)
		assert.Contains(t, out, te.out, "TestValues fail >%v", out)
	}
}
