package loop1

func bulkSend(numMessages int) float64 {
	sum := 0.0
	for i := 0; i < numMessages; i++ {

		messageCost := 1.0 + (float64(i) / 100.0)

		sum = sum + float64(messageCost)
	}

	return sum
}
