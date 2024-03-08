package bytes

func Chunk(input []byte, size int, n int) [][]byte {
	if n == 0 {
		n = len(input) / size
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
