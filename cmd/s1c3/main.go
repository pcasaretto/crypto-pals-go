package main

import (
	"encoding/hex"
	"fmt"

	"github.com/pcasaretto/crypto-pals/crypto"
)

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decoded, _ := hex.DecodeString(input)
	text, key := crypto.BreakSingleByteXOR(decoded)
	fmt.Println(string(text), string(key))
}
