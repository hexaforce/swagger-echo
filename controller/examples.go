package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hexaforce/swagger-echo/httputil"
	"github.com/labstack/echo"
)

// PingExample godoc
// @Summary ping example
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/ping [get]
func (c *Controller) PingExample(e echo.Context) {
	e.String(http.StatusOK, "pong")
	return
}

// CalcExample godoc
// @Summary calc example
// @Description plus
// @Tags example
// @Accept json
// @Produce json
// @Param val1 query int true "used for calc"
// @Param val2 query int true "used for calc"
// @Success 200 {integer} integer "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/calc [get]
func (c *Controller) CalcExample(e echo.Context) {
	val1, err := strconv.Atoi(e.QueryParam("val1"))
	if err != nil {
		httputil.NewError(e, http.StatusBadRequest, err)
		return
	}
	val2, err := strconv.Atoi(e.QueryParam("val2"))
	if err != nil {
		httputil.NewError(e, http.StatusBadRequest, err)
		return
	}
	ans := val1 + val2
	e.String(http.StatusOK, fmt.Sprintf("%d", ans))
}

// PathParamsExample godoc
// @Summary path params example
// @Description path params
// @Tags example
// @Accept json
// @Produce json
// @Param group_id path int true "Group ID"
// @Param account_id path int true "Account ID"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/groups/{group_id}/accounts/{account_id} [get]
func (c *Controller) PathParamsExample(e echo.Context) {
	groupID, err := strconv.Atoi(e.Param("group_id"))
	if err != nil {
		httputil.NewError(e, http.StatusBadRequest, err)
		return
	}
	accountID, err := strconv.Atoi(e.Param("account_id"))
	if err != nil {
		httputil.NewError(e, http.StatusBadRequest, err)
		return
	}
	e.String(http.StatusOK, fmt.Sprintf("group_id=%d account_id=%d", groupID, accountID))
}

// HeaderExample godoc
// @Summary custome header example
// @Description custome header
// @Tags example
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/header [get]
func (c *Controller) HeaderExample(e echo.Context) {
	e.String(http.StatusOK, e.Request().Header["Authorization"][0])
}

// SecuritiesExample godoc
// @Summary custome header example
// @Description custome header
// @Tags example
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Security ApiKeyAuth
// @Security OAuth2Implicit[admin, write]
// @Router /examples/securities [get]
func (c *Controller) SecuritiesExample(e echo.Context) {
}

// AttributeExample godoc
// @Summary attribute example
// @Description attribute
// @Tags example
// @Accept json
// @Produce json
// @Param enumstring query string false "string enums" Enums(A, B, C)
// @Param enumint query int false "int enums" Enums(1, 2, 3)
// @Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Param int query int false "int valid" mininum(1) maxinum(10)
// @Param default query string false "string default" default(A)
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/attribute [get]
func (c *Controller) AttributeExample(e echo.Context) {
	e.String(http.StatusOK, fmt.Sprintf("enumstring=%s enumint=%s enumnumber=%s string=%s int=%s default=%s",
		e.QueryParam("enumstring"),
		e.QueryParam("enumint"),
		e.QueryParam("enumnumber"),
		e.QueryParam("string"),
		e.QueryParam("int"),
		e.QueryParam("default"),
	))
}
