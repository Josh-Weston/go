package util

// MinInt64 returns the smallest value passed-in
func MinInt64(a ...int64) int64 {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

// MaxInt64 returns the largest value passed-in
func MaxInt64(a ...int64) int64 {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

// MinInt returns the smallest value passed-in
func MinInt(a ...int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

// MaxInt returns the largest value passed-in
func MaxInt(a ...int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}
