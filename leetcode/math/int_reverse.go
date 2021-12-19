package math

func reverse(x int32) int32 {
	const (
		MaxInt32 = 1<<31 - 1
	)

	var symbol int32 = 1
	if x < 0 {
		symbol = -1
		x = -x
	}
	var res int32 = 0
	for x > 0 && res <= MaxInt32 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > MaxInt32 {
		res = 0
	}
	return res * symbol
}
