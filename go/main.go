package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	contents, err := ioutil.ReadFile("../sample.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	benchmark := func() {
		for i := 0; i < 1e4; i++ {
			lastValue := ""
			for _, line := range lines {
				LevenshteinDistance(lastValue, line)
				lastValue = line
			}
		}
	}

	start := time.Now()
	benchmark()
	end := time.Since(start)

	fmt.Printf("%f", end.Seconds())
}

var cache [1024]int

// LevenshteinDistance determines the "edit distance" between two strings
func LevenshteinDistance(source, target string) int {
	if len(source) == 0 {
		return len(target)
	}

	if len(target) == 0 {
		return len(source)
	}

	sourceChars := []rune(source)
	targetChars := []rune(target)

	for i := 0; i < len(target)+1; i++ {
		cache[i] = i
	}

	for i, sourceChar := range sourceChars {
		nextDist := i + 1
		for j, targetChar := range targetChars {
			currentDist := nextDist

			distIfSubstitute := cache[j]
			if sourceChar != targetChar {
				distIfSubstitute++
			}

			distIfInsert := currentDist + 1
			distIfDelete := cache[j+1] + 1

			nextDist = min3(
				distIfDelete,
				distIfInsert,
				distIfSubstitute,
			)

			cache[j] = currentDist
		}

		cache[len(target)] = nextDist
	}

	return cache[len(target)]
}

func min3(a, b, c int) int {
	return min(a, min(b, c))
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
