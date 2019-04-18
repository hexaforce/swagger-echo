package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hexaforce/swagger-echo/httputil"
	"github.com/hexaforce/swagger-echo/model"
	"github.com/labstack/echo"
)

// Auth godoc
// @Summary Auth admin
// @Description get admin info
// @Tags accounts,admin
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Admin
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security ApiKeyAuth
// @Router /admin/auth [post]
func (c *Controller) Auth(e echo.Context) {
	authHeader := e.Request().Header["Authorization"][0]
	if len(authHeader) == 0 {
		httputil.NewError(e, http.StatusBadRequest, errors.New("please set Header Authorization"))
		return
	}
	if authHeader != "admin" {
		httputil.NewError(e, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to operation key=%s expected=admin", authHeader))
		return
	}
	admin := model.Admin{
		ID:   1,
		Name: "admin",
	}
	e.JSON(http.StatusOK, admin)
}
