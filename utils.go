package tcg

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sgn(a int) int {
	if a == 0 {
		return 0
	}
	if a > 0 {
		return 1
	}
	return -1
}
