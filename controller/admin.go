package controller

import (
	"errors"
	"fmt"
	"net/http"

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
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /admin/auth [post]
func (c *Controller) Auth(ctx echo.Context) error {
	authHeader := ctx.Request().Header.Get("Authorization")
	if len(authHeader) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("please set Header Authorization").Error)
	}
	if authHeader != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to operation key=%s expected=admin", authHeader).Error)
	}
	admin := model.Admin{
		ID:   1,
		Name: "admin",
	}
	return ctx.JSON(http.StatusOK, admin)
}
