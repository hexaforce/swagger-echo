package controller

import (
	"net/http"
	"strconv"

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
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /bottles/{id} [get]
func (c *Controller) ShowBottle(ctx echo.Context) error {
	id := ctx.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	bottle, err := model.BottleOne(bid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusOK, bottle)
}

// ListBottles godoc
// @Summary List bottles
// @Description get bottles
// @Tags bottles
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Bottle
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /bottles [get]
func (c *Controller) ListBottles(ctx echo.Context) error {
	bottles, err := model.BottlesAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusOK, bottles)
}
