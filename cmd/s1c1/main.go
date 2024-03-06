package main

// https://cryptopals.com/sets/1/challenges/1

import (
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
	"fmt"
)

// read from stdin a hex string and convert it to base64
func main() {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	decoder := hex.NewDecoder(os.Stdin)

	_, err := io.Copy(encoder, decoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding hex: %v\n", err)
		os.Exit(1)
	}
}
