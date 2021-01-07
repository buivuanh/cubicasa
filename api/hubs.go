package api

import (
	"github.com/buivuanh/cubicasa/model"
	"github.com/labstack/echo"
	"net/http"
)

func (a *API) InitHub() {
	a.Routes.Hubs.POST("",a.CreateHub)
	a.Routes.Hub.GET("",a.GetHub)
}

func (a API) CreateHub(c echo.Context) error {
	h := model.Hub{}
	err := c.Bind(&h)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := a.App.CreateHub(&h)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (a API) GetHub(c echo.Context) error {
	hubID := paramFromContext(c).HubID
	res, err := a.App.GetHub(hubID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}