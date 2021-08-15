package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/go-zelus/api-layout/entity/state"
)

type Controller struct{}

func (c *Controller) JSON(ctx *gin.Context, args interface{}) {
	if e, ok := args.(error); ok {
		switch err := e.(type) {
		case *mysql.MySQLError:
			ctx.SecureJSON(state.DatabaseError.StatusCode(), state.DatabaseError.Response(err))
			return
		case state.ValidErrors:
			ctx.SecureJSON(state.InvalidParams.StatusCode(), state.InvalidParams.Response(err))
			return
		case *state.Error:
			ctx.JSON(err.StatusCode(), err.Response())
			return
		default:
			ne := state.ServerError
			ctx.JSON(ne.StatusCode(), ne.Response(e))
			return
		}
	}
	ctx.SecureJSON(state.Success.StatusCode(), args)
}

func (c *Controller) BindAndValid(ctx *gin.Context, v interface{}) (bool, state.ValidErrors) {
	var errs state.ValidErrors
	err := ctx.ShouldBind(v)
	fmt.Println(err)
	if err != nil {
		val := ctx.Value("trans")
		trans, _ := val.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &state.ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
