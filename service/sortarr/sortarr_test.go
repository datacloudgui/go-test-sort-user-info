// Package sortarray_test provide a simple test for the sorted method
package sortarray_test

import (
	"errors"
	sortarray "main/service/sortarr"
	"testing"
)

//TestSortArr
func TestSortArr(t *testing.T) {

	var arr = []int{3, 5, 5, 6, 8, 3, 4, 4, 7, 7, 1, 1, 2}
	var arrSort = []int{1, 2, 3, 4, 5, 6, 7, 8, 5, 3, 4, 7, 1}

	t.Log("Trying to sort: ", arr)

	sorted := sortarray.SortArrDuplicatesEnd(arr)

	for index, number := range sorted {
		if number != arrSort[index] {
			t.Fatalf("Fail test sortArr - %v", errors.New("Unexpected order"))
		}
	}

	t.Log("Array sorted OK")
}
