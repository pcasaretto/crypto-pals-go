package crypto_test

import (
	"testing"

	"github.com/pcasaretto/crypto-pals/crypto"
)

func TestFixedXOR(t *testing.T) {
	expected := "746865206b696420646f6e277420706c6179"
	actual := crypto.FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestBreakSingleByteXOR(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	actual := crypto.BreakSingleByteXOR(input)
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestEncryptRepeatingXOR(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	actual := crypto.EncryptRepeatingXOR(input, "ICE")
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
