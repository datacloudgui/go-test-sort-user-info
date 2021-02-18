package sort

import (
	"encoding/json"
	"fmt"
	"main/tools"

	"github.com/valyala/fasthttp"
)

// Unsorted ..
type Unsorted struct {
	Unsorted []int
}

// SortArray ...
func SortArray(ctx *fasthttp.RequestCtx) {

	fmt.Println("Sort the array: ")

	arr := &Unsorted{}

	err := json.Unmarshal(ctx.PostBody(), arr)
	if err != nil {
		tools.ResponseHandlers(ctx, nil, "Error json decode: the body request is null or invalid datatype", fasthttp.StatusUnprocessableEntity)
		return
	}

	fmt.Println("El contenido de request es: ", arr.Unsorted[0])

	// response, err := sort(arr.Unsorted)
	// if err != nil {
	// 	tools.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError)
	// 	return
	// }

	tools.ResponseHandlers(ctx, nil, nil, fasthttp.StatusOK)

}
