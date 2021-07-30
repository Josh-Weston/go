package util

// Min returns the smallest value in an integer slice.
// The runtime will panic if slice is empty; therefore, the user should ensure the slice is populated
func Min(a []int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

// Max returns the smallest value in an integer slice.
// The runtime will panic if slice is empty; therefore, the user should ensure the slice is populated
func Max(a []int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

// Extent returns the min and max values from the array
// The runtime will panic if slice is empty; therefore, the user should ensure the slice is populated
func Extent(a []int) (int, int) {
	min := a[0]
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}
