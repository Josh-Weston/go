package util

// Equal compares the elements of two slices by first checking if they are the same length, and then comparing
// each element individually
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Extent returns the min and max values in an int slice
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

// Reverse reverses the slice in place
func Reverse() {
	/*
		for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
			a[left], a[right] = a[right], a[left]
		}
	*/
}

func Splice(a *[]int, i int, deleteCount int, items ...int) []int {
	arr := *a
	if len(arr) == 0 {
		*a = append(arr, items...)
		return []int{}
	}
	if deleteCount <= 0 {
		*a = append(arr[:i], append(items, arr[i:]...)...)
		return []int{}
	}
	d := append(make([]int, 0, len(arr[i:i+deleteCount])), arr[i:i+deleteCount]...)
	*a = append(arr[:i], append(items, arr[i+deleteCount:]...)...)
	return d
}

// Pop removes the last element from the slice
func Pop() {

}

// Shift removes the first element from the slice
func Shift() {

}

// Unshift adds elements to the start of the slice
func Unshift() {

}
