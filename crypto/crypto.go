package crypto

import (
	"encoding/hex"
)

func EncryptRepeatingXOR(input, key []byte) []byte {
	inputBytes := []byte(input)
	copy(inputBytes, input)

	// XOR the input with the key
	for i := range inputBytes {
		inputBytes[i] ^= key[i%len(key)]
	}

	return inputBytes
}

func DecryptRepeatingXOR(input, key []byte) []byte {
	// make a copy of the input
	inputBytes := make([]byte, len(input))
	copy(inputBytes, input)
	// XOR the input with the key
	for i := range inputBytes {
		inputBytes[i] ^= key[i%len(key)]
	}

	return inputBytes
}

func BreakSingleByteXOR(input []byte) ([]byte, byte) {
	// make a copy of the input
	bytes := make([]byte, len(input))
	copy(bytes, input)
	// score each byte
	var bestScore int
	var bestByte byte
	for b := 0; b < 256; b++ {
		score := scoreXOR(input, byte(b))
		if score > bestScore {
			bestScore = score
			bestByte = byte(b)
		}
	}

	// XOR the bytes with the best byte
	for i := range bytes {
		bytes[i] ^= bestByte
	}

	return bytes, bestByte
}

func scoreXOR(bytes []byte, b byte) int {
	var score int
	for _, c := range bytes {
		score += scoreChar(c ^ b)
	}
	return score
}

func scoreChar(c byte) int {
	// score the character
	switch {
	case c >= 'a' && c <= 'z':
		return 1
	case c >= 'A' && c <= 'Z':
		return 1
	case c == ' ':
		return 1
	case c == '\n':
		return 1
	case c == '\t':
		return 1
	default:
		return 0
	}
}

// FixedXOR takes two equal-length buffers and produces their XOR combination
func FixedXOR(hex1, hex2 string) string {
	if len(hex1) != len(hex2) {
		return ""
	}

	// convert hex to bytes
	bytes1, err := hex.DecodeString(hex1)
	if err != nil {
		return ""
	}
	bytes2, err := hex.DecodeString(hex2)
	if err != nil {
		return ""
	}

	// XOR the bytes
	for i := range bytes1 {
		bytes1[i] ^= bytes2[i]
	}

	// convert bytes to hex
	return hex.EncodeToString(bytes1)
}

func HammingDistance(b1, b2 []byte) int {
	// XOR the bytes
	var distance int
	for i := range b1 {
		distance += popCount(b1[i] ^ b2[i])
	}

	return distance
}

func popCount(b byte) int {
	// count the number of bits set in a byte
	var count int
	for b != 0 {
		count++
		b &= b - 1
	}
	return count
}
