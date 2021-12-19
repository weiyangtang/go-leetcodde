package hash

func isHappy(n int) bool {
	hashMap := make(map[int]bool, 0)
	isFind, _ := hashMap[n]
	for !isFind && n != 1 {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		isFind, _ = hashMap[sum]
		n = sum
	}

	if n == 1 {
		return true
	}
	return false
}
