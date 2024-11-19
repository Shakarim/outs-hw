package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(sourceValue string) []string {
	// Place your code here.
	var words []string = strings.Fields(sourceValue)
	var numberOfDash = strings.Count(sourceValue, "-")
	var result map[string]int = make(map[string]int, len(words)+(numberOfDash*2))

	for _, word := range words {
		indicateWordInMap(result, word)
	}

	keys := make([]string, 0, len(result))

	for key := range result {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return result[keys[i]] > result[keys[j]]
	})

	if len(keys) >= 10 {
		return keys[:10]
	}

	return keys
}

func indicateWordInMap(result map[string]int, word string) {
	if word != "" {
		val, ok := result[word]
		if ok {
			result[word] = val + 1
		} else {
			result[word] = 1
		}
	}
}
