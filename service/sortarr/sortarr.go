package sortarray

import (
	"encoding/json"
	"main/tools"
	"sort"

	"github.com/valyala/fasthttp"
)

//Structs for input and output JSON data

// Unsorted ..
type Unsorted struct {
	Unsorted []int
}

// SortedResponse ..
type SortedResponse struct {
	Unsorted []int
	Sorted   []int
}

// SortArray ...
func SortArray(ctx *fasthttp.RequestCtx) {

	arr := &Unsorted{}

	err := json.Unmarshal(ctx.PostBody(), arr)
	if err != nil {
		tools.ResponseHandlers(ctx, nil, "Error json decode: the body request is null or invalid datatype", fasthttp.StatusUnprocessableEntity)
		return
	}

	if len(arr.Unsorted) < 2 {
		tools.ResponseHandlers(ctx, nil, "Error array must be contain more than 1 value", fasthttp.StatusUnprocessableEntity)
		return
	}

	sorted := &SortedResponse{}

	sorted.Unsorted = arr.Unsorted
	sorted.Sorted = SortArrDuplicatesEnd(arr.Unsorted)

	tools.ResponseHandlers(ctx, sorted, nil, fasthttp.StatusOK)

}

// Separate duplicated values and sort unique values
// SortArrDuplicatesEnd ..
func SortArrDuplicatesEnd(numbers []int) []int {

	var duplicated []int
	var sorted []int

	for _, number := range numbers {
		_, found := find(sorted, number)
		if !found {
			sorted = append(sorted, number)
		} else {
			duplicated = append(duplicated, number)
		}
	}

	//Sort the array with unique values
	sort.Ints(sorted)

	return append(sorted, duplicated...)
}

// find if a value is present in a passed slice
func find(numbers []int, val int) (int, bool) {
	for index, item := range numbers {
		if item == val {
			return index, true
		}
	}
	return -1, false
}
