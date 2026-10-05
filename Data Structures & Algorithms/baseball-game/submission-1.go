

func calPoints(operations []string) int {
	init := false
	lastItem := 0
	scores := make([]int, len(operations)) // the stack
	for _, op := range operations {
		v, err := strconv.Atoi(op)
		if err != nil {
			// it must be a string
			if op == "C" {
				// swap remove last item in scores
				lastItem--
			}
			if op == "D" {
				// add new item to scroes, double the last item
				double := scores[lastItem] * 2

				// add to stack
				lastItem++
				scores[lastItem] = double

			}
			if op == "+" {
				// add new item to scores, the sum of curr last and prior
				last := scores[lastItem]
				prior := scores[lastItem-1]

				// add to stack
				lastItem++
				scores[lastItem] = last + prior

			}
		} else {
			// add the value to scores
			if !init {
				scores[lastItem] = v
				init = true
			} else {
				lastItem++
				scores[lastItem] = v
			}

		}
	}

	// read over scores
	points := 0
	for i := lastItem; i >= 0; i-- {
		points = points + scores[i]
	}

	return points
}
