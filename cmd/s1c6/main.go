// https://cryptopals.com/sets/1/challenges/6
package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"os"
	"sort"

	"golang.org/x/exp/maps"

	"github.com/pcasaretto/crypto-pals/crypto"
)

type CandidateKey struct {
	Normalized float64
	Key        []byte
	Length     int
}

func main() {
	decoder := base64.NewDecoder(base64.StdEncoding, os.Stdin)
	input, err := io.ReadAll(decoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	keys := make([]CandidateKey, 0)

	for keysize := 2; keysize <= 40; keysize++ {
		chunks := chunk(input, keysize, 4)
		normalized := normalize(chunks)
		keys = append(keys, CandidateKey{Length: keysize, Normalized: normalized})
	}

	// sort keys by normalized distance
	// take the top 5
	// for each key, break the repeating xor

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Normalized < keys[j].Normalized
	})

	minScore := math.MaxFloat64
	var decrypted []byte

	for _, key := range keys[:5] {
		chunks := chunk(input, key.Length, 0)
		transposed := transpose(chunks)

		key := make([]byte, key.Length)
		for i := 0; i < len(transposed); i++ {
			_, bestByte := crypto.BreakSingleByteXOR(transposed[i])
			key[i] = bestByte
		}

		// decrypt the input using the key
		decryptedCandidate := crypto.DecryptRepeatingXOR(input, key)

		// perform statistical analysis on the decrypted text
		score := scoreText(string(decryptedCandidate))

		if score < minScore {
			minScore = score
			decrypted = decryptedCandidate
		}
	}

	fmt.Println(string(decrypted))
}

var englishFrequency = map[byte]float64{
	'a': 0.08167,
	'b': 0.01492,
	'c': 0.02782,
	'd': 0.04253,
	'e': 0.12702,
	'f': 0.02228,
	'g': 0.02015,
	'h': 0.06094,
	'i': 0.06966,
	'j': 0.00153,
	'k': 0.00772,
	'l': 0.04025,
	'm': 0.02406,
	'n': 0.06749,
	'o': 0.07507,
	'p': 0.01929,
	'q': 0.00095,
	'r': 0.05987,
	's': 0.06327,
	't': 0.09056,
	'u': 0.02758,
	'v': 0.00978,
	'w': 0.02360,
	'x': 0.00150,
	'y': 0.01974,
	'z': 0.00074,
}

// scoreText returns a score for the text using the relative
// frequency of english characters
func scoreText(text string) float64 {
	// calculate the frequency of each character
	frequency := make(map[byte]float64)
	for _, c := range text {
		frequency[byte(c)]++
	}
	// normalize the frequency
	for k := range frequency {
		frequency[k] /= float64(len(text))
	}

	return kolmogorovSmirnovishTest(frequency, englishFrequency)
}

// calculateCumulativeSums calculates the cumulative sum for each unique key in sorted order.
func calculateCumulativeSums(data map[byte]float64) []float64 {
	keys := maps.Keys(englishFrequency)

	// Calculate cumulative sums
	cumulativeSums := make([]float64, len(keys))
	var sum float64
	for i, k := range keys {
		sum += data[k]
		cumulativeSums[i] = sum
	}

	return cumulativeSums
}

// kolmogorovSmirnovTest calculates the K-S Test score between two distributions.
func kolmogorovSmirnovishTest(dist1, dist2 map[byte]float64) float64 {
	cumulative1 := calculateCumulativeSums(dist1)
	cumulative2 := calculateCumulativeSums(dist2)

	// Find the maximum difference
	totalDiff := 0.0
	for i := 0; i < len(cumulative1) && i < len(cumulative2); i++ {
		diff := abs(cumulative1[i] - cumulative2[i])
		totalDiff += diff
	}

	return totalDiff / float64(len(cumulative1))
}

// abs returns the absolute value of a float64.
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func transpose(chunks [][]byte) [][]byte {
	transposed := make([][]byte, len(chunks[0]))
	for i := 0; i < len(chunks[0]); i++ {
		transposed[i] = make([]byte, 0, len(chunks))
		for j := 0; j < len(chunks); j++ {
			if i >= len(chunks[j]) {
				continue
			}
			transposed[i] = append(transposed[i], chunks[j][i])
		}
	}
	return transposed
}

func chunk(input []byte, size int, n int) [][]byte {
	if n == 0 {
		n = len(input)/size + 1
	}
	chunks := make([][]byte, 0, len(input)/size)
	for i := 0; i < n; i++ {
		start := i * size
		end := start + size
		if end > len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[start:end])
	}
	return chunks
}

func normalize(chunks [][]byte) float64 {
	var total float64
	for i := 0; i < len(chunks)-1; i++ {
		total += float64(crypto.HammingDistance(chunks[i], chunks[i+1])) / float64(len(chunks[i]))
	}
	return total / float64(len(chunks)-1)
}
