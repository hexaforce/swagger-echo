package controller

import (
	"net/http"
	"strconv"

	"github.com/hexaforce/swagger-echo/httputil"
	"github.com/hexaforce/swagger-echo/model"
	"github.com/labstack/echo"
)

// ShowBottle godoc
// @Summary Show a bottle
// @Description get string by ID
// @ID get-string-by-int
// @Tags bottles
// @Accept  json
// @Produce  json
// @Param  id path int true "Bottle ID"
// @Success 200 {object} model.Bottle
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /bottles/{id} [get]
func (c *Controller) ShowBottle(e echo.Context) {
	id := e.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(e, http.StatusBadRequest, err)
		return
	}
	bottle, err := model.BottleOne(bid)
	if err != nil {
		httputil.NewError(e, http.StatusNotFound, err)
		return
	}
	e.JSON(http.StatusOK, bottle)
}

// ListBottles godoc
// @Summary List bottles
// @Description get bottles
// @Tags bottles
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Bottle
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /bottles [get]
func (c *Controller) ListBottles(e echo.Context) {
	bottles, err := model.BottlesAll()
	if err != nil {
		httputil.NewError(e, http.StatusNotFound, err)
		return
	}
	e.JSON(http.StatusOK, bottles)
}
