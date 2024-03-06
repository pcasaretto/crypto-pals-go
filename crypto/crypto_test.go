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
