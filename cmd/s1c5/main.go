package main

import (
	"fmt"

	"github.com/pcasaretto/crypto-pals/crypto"
)

func main() {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	fmt.Println(crypto.EncryptRepeatingXOR(input, "ICE"))
}
