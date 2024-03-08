package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/pcasaretto/crypto-pals/bytes"
	"github.com/pcasaretto/crypto-pals/crypto"
)

type Candidate struct {
	LineNumber int
	Similarity float64
	Text       []byte
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	lineNumber := 0
	best := Candidate{}
	for scanner.Scan() {
		line := scanner.Text()
		decoded, err := hex.DecodeString(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding hex: %v\n", err)
			os.Exit(1)
		}
		candidate := Candidate{LineNumber: lineNumber, Similarity: calculateSimilarity(decoded), Text: decoded}
		if candidate.Similarity > best.Similarity {
			best = candidate
		}
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Best candidate: %d\n", best.LineNumber)

	fmt.Println(hex.EncodeToString(best.Text))
}

// https://www3.nd.edu/~busiforc/handouts/cryptography/Letter%20Frequencies.html#quadrigrams
var commonEnglish16ByteBlocks = map[string]float64{
	"that": 0.761242,
	"ther": 0.604501,
	"with": 0.573866,
	"tion": 0.551919,
	"here": 0.374549,
	"ould": 0.369920,
	"ight": 0.309440,
	"have": 0.290544,
	"hich": 0.284292,
	"whic": 0.283826,
	"this": 0.276333,
	"thin": 0.270413,
	"they": 0.262421,
	"atio": 0.262386,
	"ever": 0.260695,
	"from": 0.258580,
	"ough": 0.253447,
	"were": 0.231089,
	"hing": 0.229944,
	"ment": 0.223347,
}

func calculateSimilarity(input []byte) float64 {
	// split input into 16 byte blocks
	// for each block, calculate the hamming distance with the next block
	// sum all the distances and divide by the number of blocks
	// return the average distance

	// split input into 16 byte blocks
	blocks := bytes.Chunk(input, 16, 0)
	equalBlocks := 0

	for i := 0; i < len(blocks)-1; i++ {
		for j := i + 1; j < len(blocks)-1; j++ {
			if i == j {
				continue
			}
			// calculate the hamming distance with the next block
			if crypto.HammingDistance(blocks[i], blocks[j]) == 0 {
				equalBlocks++
			}
		}
	}
	return float64(equalBlocks) / float64(len(blocks))
}
