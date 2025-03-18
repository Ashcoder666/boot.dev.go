package sli12

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for index, word := range msg {
		for _, w := range badWords {
			if w == word {
				return index
			}
		}
	}
	return -1
}
