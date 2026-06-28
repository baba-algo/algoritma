var stack []int

func push(value int) {
	stack = append(stack, value)
}

func pop() (int, bool) {
	if len(stack) == 0 {
		return 0, false
	}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return top, true
}

func peek() (int, bool) {
	if len(stack) == 0 {
		return 0, false
	}
	return stack[len(stack)-1], true
}
