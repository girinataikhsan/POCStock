package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Show(ctx echo.Context, httpstatus int, code int, msg string, data interface{}) error {
	ctx.Response().Header().Add("Strict-Transport-Security", "max-age=0; includeSubDomains")
	ctx.Response().Header().Add("X-Content-Type-Options", "nosniff")

	return ctx.JSON(httpstatus, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
