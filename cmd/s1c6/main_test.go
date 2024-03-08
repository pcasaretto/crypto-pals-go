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
