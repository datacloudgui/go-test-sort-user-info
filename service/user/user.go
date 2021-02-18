package user

import (
	"fmt"
	"main/orm"
	"main/tools"

	"github.com/valyala/fasthttp"
)

// User ..
type User struct {
	ID          string
	CompanyName string
	RoleName    string
	Name        string
	LastName    string
	Email       string
	IsActive    bool
}

// GetUser ...
func GetUser(ctx *fasthttp.RequestCtx) {

	fmt.Println("Getting user")

	id, ok := ctx.UserValue("id").(string)
	if !ok {
		tools.ResponseHandlers(ctx, nil, "Invalid ID request", fasthttp.StatusUnprocessableEntity)
		return
	}

	if id == "" {
		tools.ResponseHandlers(ctx, nil, "ID cannot be empty", fasthttp.StatusUnprocessableEntity)
		return
	}

	ormResponse, err := orm.GetUser(id)
	if err != nil {
		tools.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError)
		return
	}

	tools.ResponseHandlers(ctx, ormResponse, nil, fasthttp.StatusOK)

}
