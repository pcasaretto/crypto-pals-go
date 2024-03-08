// https://cryptopals.com/sets/1/challenges/7
package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	decoder := base64.NewDecoder(base64.StdEncoding, os.Stdin)
	input, err := io.ReadAll(decoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}
	decoded := DecryptAes128Ecb(input, []byte("YELLOW SUBMARINE"))
	fmt.Println(string(decoded))
}

func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
