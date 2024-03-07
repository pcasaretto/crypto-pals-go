package crypto_test

import (
	"encoding/hex"
	"testing"

	"github.com/pcasaretto/crypto-pals/crypto"
	"github.com/stretchr/testify/assert"
)

func TestFixedXOR(t *testing.T) {
	expected := "746865206b696420646f6e277420706c6179"
	actual := crypto.FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	assert.Equal(t, expected, actual)
}

func TestBreakSingleByteXOR(t *testing.T) {
	input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	expected := []byte("Cooking MC's like a pound of bacon")
	actual, _ := crypto.BreakSingleByteXOR(input)
	assert.Equal(t, expected, actual)
}

func TestDecryptRepeatingXOR(t *testing.T) {
	// test that any input string can be encrypted and decrypted
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	encrypted := crypto.EncryptRepeatingXOR([]byte(input), []byte(key))
	decrypted := string(crypto.DecryptRepeatingXOR([]byte(encrypted), []byte(key)))
	assert.Equal(t, input, decrypted)
}

func TestHammingDistance(t *testing.T) {
	expected := 37
	actual := crypto.HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	assert.Equal(t, expected, actual)
}
