package sli11

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	newSlice := []float64{}
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			newSlice = append(newSlice, costs[i].value)
		}
	}
	return newSlice

}
