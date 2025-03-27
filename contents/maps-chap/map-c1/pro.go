package mapc1

import "strings"

func countDistinctWords(messages []string) int {
	distinctiveMap := make(map[string]int)
	for _, message := range messages {
		wordsArr := strings.Fields(message)

		for _, word := range wordsArr {
			wordLowered := strings.ToLower(word)
			distinctiveMap[wordLowered] = 1
		}
	}

	return len(distinctiveMap)
}
