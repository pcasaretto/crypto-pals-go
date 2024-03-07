package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	input := [][]byte{
		[]byte("abc"),
		[]byte("def"),
		[]byte("ghi"),
	}
	expected := [][]byte{
		[]byte("adg"),
		[]byte("beh"),
		[]byte("cfi"),
	}
	actual := transpose(input)
	assert.Equal(t, expected, actual)
}

func TestTranspose2(t *testing.T) {
	input := [][]byte{
		[]byte("abc"),
		[]byte("def"),
		[]byte("g"),
	}
	expected := [][]byte{
		[]byte("adg"),
		[]byte("be"),
		[]byte("cf"),
	}
	actual := transpose(input)
	assert.Equal(t, expected, actual, "expected: %s, actual: %s", expected, actual)
}

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
		actual := chunk(test.input, test.size, test.n)
		assert.Equal(t, test.expected, actual, "test %d: expected: %s, actual: %s", k, test.expected, actual)
	}
}
