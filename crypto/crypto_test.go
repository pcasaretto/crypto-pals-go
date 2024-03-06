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
