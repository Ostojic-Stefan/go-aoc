package utils

func IntMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
