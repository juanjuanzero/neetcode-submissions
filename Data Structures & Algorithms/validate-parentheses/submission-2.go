// a stack problem


const (
	paren   rune = ')'
	bracket rune = '}'
	square  rune = ']'
)

func isValid(s string) bool {
	lastIndex := 0
	stack := make([]rune, len(s))
	// parenthesis will work since they are ascii characters
	for i, v := range s {
		if i != 0 && lastIndex != -1 {
			// check if the last item in the stack
			lastItem := stack[lastIndex]
			if lastItem == '(' && v == paren {
				lastIndex--
			} else if lastItem == '[' && v == square {
				lastIndex--
			} else if lastItem == '{' && v == bracket {
				lastIndex--
			} else {
				// add it
				lastIndex++
				stack[lastIndex] = v
			}
		} else {
			// insert the first item
			lastIndex = 0
			stack[lastIndex] = v
		}
	}
	// only valid if lastIndex as decremented to -1
	return lastIndex == -1
	}
