package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode/utf8"
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

	// check
	dists := make([]string, len(lines)-1)
	for i := 0; i < len(dists); i++ {
		dist := LevenshteinDistance(lines[i], lines[i+1])
		dists[i] = fmt.Sprintf("%d", dist)
	}
	fmt.Fprintln(os.Stderr, strings.Join(dists, ","))
}

// LevenshteinDistance determines the "edit distance" between two strings
func LevenshteinDistance(source, target string) int {
	if len(source) == 0 {
		return utf8.RuneCountInString(target)
	}

	if len(target) == 0 {
		return utf8.RuneCountInString(source)
	}

	targetLen := utf8.RuneCountInString(target)
	cache := make([]int, targetLen+1)
	for i := 0; i < targetLen+1; i++ {
		cache[i] = i
	}

	for i, sourceChar := range source {
		nextDist := i + 1
		for j, targetChar := range target {
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

		cache[targetLen] = nextDist
	}

	return cache[targetLen]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
