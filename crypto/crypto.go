package crypto

import (
	"encoding/hex"
)

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
