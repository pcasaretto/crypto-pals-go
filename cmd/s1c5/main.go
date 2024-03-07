package main

import (
	"encoding/hex"
	"fmt"

	"github.com/pcasaretto/crypto-pals/crypto"
)

func main() {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	text := crypto.EncryptRepeatingXOR([]byte(input), []byte("ICE"))
	encoded := hex.EncodeToString(text)
	fmt.Println(encoded)
}
