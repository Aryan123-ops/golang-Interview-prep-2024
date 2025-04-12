package main

import (
	"fmt"
	"sort"
)

type charCount struct {
	char  rune
	count int
}

func main() {
	s := "aabbbbcccddddd"
	freq := make(map[rune]int)

	// Count frequency of each character
	for _, ch := range s {
		freq[ch]++
	}

	// Convert map to slice for sorting
	var counts []charCount
	for ch, cnt := range freq {
		counts = append(counts, charCount{ch, cnt})
	}

	// Sort slice by count in descending order
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	// Build result string
	result := ""
	for _, c := range counts {
		result += fmt.Sprintf("%d%c", c.count, c.char)
	}

	fmt.Println(result)
}
