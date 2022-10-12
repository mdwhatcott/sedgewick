package ch1p5

import (
	"os"
	"strconv"
	"strings"
)

var (
	tiny   = readInts("tinyUF.txt")
	medium = readInts("mediumUF.txt")
	large  = readInts("largeUF.txt")
)

func parseInt(v string) int {
	n, _ := strconv.Atoi(v)
	return n
}
func readInts(filename string) (results []int) {
	content, _ := os.ReadFile(filename)
	values := strings.Fields(string(content))
	results = make([]int, len(values))
	for r, raw := range values {
		results[r] = parseInt(raw)
	}
	return results
}
