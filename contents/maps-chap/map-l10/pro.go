package mapl10

func getNameCounts(names []string) map[rune]map[string]int {
	database := make(map[rune]map[string]int)
	databaseContent := make(map[string]int)
	for _, word := range names {
		firstLetter := rune(word[0])
		databaseContent[word]++
		database[firstLetter] = databaseContent

	}

	return database
}
