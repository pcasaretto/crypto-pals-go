package main

import (
	"fmt"
	"github.com/pcasaretto/crypto-pals/crypto"
)


func main() {
	fmt.Println(crypto.FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
}
