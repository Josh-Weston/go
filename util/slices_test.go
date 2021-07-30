package util

import (
	"testing"
)

func TestSplice(t *testing.T) {

	// Test with empty slice, no insertion
	result := Splice(&[]int{}, 0, 0)
	if len(result) != 0 {
		t.Errorf("TestSplice (empty slice, no insertion) failed ")
	}

	// Test with empty slice, with insertion
	arr := []int{}
	result = Splice(&arr, 5, 0, []int{10, 20, 30}...)
	if !Equal(result, []int{}) || !Equal(arr, []int{10, 20, 30}) {
		t.Errorf("TestSplice (empty slice, with insertion) failed ")
	}

	// Test insertion with no delete count
	arr = []int{1, 2, 3, 4, 5}
	result = Splice(&arr, 2, 0, []int{6, 7}...)
	if !Equal(result, []int{}) || !Equal(arr, []int{1, 2, 6, 7, 3, 4, 5}) {
		t.Errorf("TestSplice (insertion, no delete) failed ")
	}

	// Test delete with no insertion
	arr = []int{1, 2, 3, 4, 5}
	result = Splice(&arr, 2, 2)
	if !Equal(result, []int{3, 4}) || !Equal(arr, []int{1, 2, 5}) {
		t.Errorf("TestSplice (delete, no insertion) failed ")
	}

	// Test with delete with insertion
	arr = []int{1, 2, 3, 4, 5}
	result = Splice(&arr, 2, 2, []int{6, 7}...)
	if !Equal(result, []int{3, 4}) || !Equal(arr, []int{1, 2, 6, 7, 5}) {
		t.Errorf("TestSplice (delete and insertion) failed ")
	}

	// Test with negative delete with insertion
	arr = []int{1, 2, 3, 4, 5}
	result = Splice(&arr, 2, -2, []int{6, 7}...)
	if !Equal(result, []int{}) || !Equal(arr, []int{1, 2, 6, 7, 3, 4, 5}) {
		t.Errorf("TestSplice (negative delete and insertion) failed ")
	}

}

func BenchmarkSplice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Splice(&[]int{1, 2, 3, 4, 5}, 2, 2, []int{6, 7}...)
	}
}
