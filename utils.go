package tcg

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
