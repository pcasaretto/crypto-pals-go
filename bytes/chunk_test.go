package bytes_test

import (
	"testing"

	"github.com/pcasaretto/crypto-pals/bytes"
	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {

	tests := []struct {
		input    []byte
		size     int
		n        int
		expected [][]byte
	}{
		{
			[]byte("abcdefghij"),
			4,
			0,
			[][]byte{
				[]byte("abcd"),
				[]byte("efgh"),
				[]byte("ij"),
			},
		},
		{
			[]byte("abcdefghij"),
			4,
			2,
			[][]byte{
				[]byte("abcd"),
				[]byte("efgh"),
			},
		},
		{
			[]byte("abcdefghij"),
			4,
			3,
			[][]byte{
				[]byte("abcd"),
				[]byte("efgh"),
				[]byte("ij"),
			},
		},
	}

	for k, test := range tests {
		actual := bytes.Chunk(test.input, test.size, test.n)
		assert.Equal(t, test.expected, actual, "test %d: expected: %s, actual: %s", k, test.expected, actual)
	}
}
