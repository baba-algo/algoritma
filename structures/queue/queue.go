var queue []int

func enqueue(value int) {
	queue = append(queue, value)
}

func dequeue() (int, bool) {
	if len(queue) == 0 {
		return 0, false
	}
	front := queue[0]
	queue = queue[1:]
	return front, true
}
