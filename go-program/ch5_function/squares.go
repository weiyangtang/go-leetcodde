package ch5_function

func squares() func() int {
	var x int
	return func() int {
		x++
		return x + x
	}
}
