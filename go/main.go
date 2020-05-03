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

	cache := make([]int, len(target)+1)
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

			nextDist = min(distIfDelete, min(distIfInsert, distIfSubstitute))

			cache[j] = currentDist
		}

		cache[len(target)] = nextDist
	}

	return cache[len(target)]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
