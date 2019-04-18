package httputil

import "github.com/labstack/echo"

// NewError example
func NewError(e echo.Context, status int, err error) {
	er := echo.HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	e.JSON(status, er)
}

// HTTPError example
// type HTTPError struct {
// 	Code    int    `json:"code" example:"400"`
// 	Message string `json:"message" example:"status bad request"`
// }
