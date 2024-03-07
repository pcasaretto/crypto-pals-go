package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/pcasaretto/crypto-pals/crypto"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	candidates := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		decoded, err := hex.DecodeString(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding hex: %v\n", err)
			os.Exit(1)
		}
		candidate, _ := crypto.BreakSingleByteXOR(decoded)
		candidates = append(candidates, string(candidate))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	best := bestCandidate(candidates)
	fmt.Println(best)
}

func bestCandidate(candidates []string) string {
	best := ""
	bestScore := 0
	for _, candidate := range candidates {
		score := frequencyAnalysis(candidate)
		if score > bestScore {
			best = candidate
			bestScore = score
		}
	}
	return best
}

func frequencyAnalysis(candidate string) int {
	score := 0
	for _, c := range candidate {
		if strings.Contains("etaoinshrdlu", string(c)) {
			score++
		}
	}
	return score
}
