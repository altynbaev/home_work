package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	inputSlice := strings.Fields(s)

	wordsMap := map[string]int{}
	wordsSlice := make([]string, 0, len(inputSlice))

	for _, word := range inputSlice {
		wordsMap[word]++
	}

	for word := range wordsMap {
		wordsSlice = append(wordsSlice, word)
	}

	sort.Slice(wordsSlice, func(i int, j int) bool {
		// Условие для лексикографической сортировки при равном количестве слов
		if wordsMap[wordsSlice[i]] == wordsMap[wordsSlice[j]] {
			return wordsSlice[i] < wordsSlice[j]
		}
		// Условие для сортировки по количеству
		return wordsMap[wordsSlice[i]] > wordsMap[wordsSlice[j]]
	})

	var sliceLen int
	if len(inputSlice) < 10 {
		sliceLen = len(wordsSlice)
	} else {
		sliceLen = 10
	}

	return wordsSlice[:sliceLen]
}
