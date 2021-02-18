package tools

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// RestErrorResponse ...
type RestErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// ResponseHandlers ...
func ResponseHandlers(ctx *fasthttp.RequestCtx, data interface{}, err interface{}, statusCode int) {

	ctx.SetContentType("application/json; charset=UTF-8")
	ctx.SetStatusCode(statusCode)

	str := struct {
		Response interface{} `json:"response"`
		Error    interface{} `json:"error"`
	}{
		Response: data,
		Error:    err,
	}

	serialized, err := json.Marshal(str)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte("Serialization Error"))
		return
	}

	ctx.SetBody(serialized)
}
