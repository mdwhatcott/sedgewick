package ch1p4

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkThreeSumNaive1Kints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeSumNaive(loadInts(b, "1Kints.txt"))
	}
}

func BenchmarkThreeSumNaive2Kints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeSumNaive(loadInts(b, "2Kints.txt"))
	}
}

func loadInts(b *testing.B, filename string) (numbers []int) {
	b.StopTimer()
	defer b.StartTimer()

	raw, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}

	content := strings.TrimSpace(string(raw))
	for _, line := range strings.Split(content, "\n") {
		text := strings.TrimSpace(line)
		n, err := strconv.Atoi(text)
		if err != nil {
			b.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}
