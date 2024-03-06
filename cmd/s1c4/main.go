package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pcasaretto/crypto-pals/crypto"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanLines)

	candidates := make([]string, 0, len(input))
	for scanner.Scan() {
		line := scanner.Text()
		candidates = append(candidates, crypto.BreakSingleByteXOR(line))
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
