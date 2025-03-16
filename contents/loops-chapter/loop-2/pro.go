package loop2

func maxMessages(thresh int) int {
	totalCost := 0
	count := 0

	for {
		cost := 100 + count

		if cost+totalCost > thresh {
			break
		}

		totalCost += cost

		count++
	}

	return count
}
