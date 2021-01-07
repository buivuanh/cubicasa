package api

import (
	"github.com/buivuanh/cubicasa/model"
	"github.com/labstack/echo"
	"net/http"
)

func (a *API) InitProductionTeam() {
	a.Routes.ProdTeams.POST("",a.CreateProductionTeam)
	a.Routes.ProdTeam.GET("",a.GetProductionTeam)
	a.Routes.ProdTeam.POST("/join-hub",a.ProductionTeamJoinHub)
}

func (a API) CreateProductionTeam(c echo.Context) error {
	pt := model.ProductionTeam{}
	err := c.Bind(&pt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := a.App.CreateProductionTeam(&pt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (a API) GetProductionTeam(c echo.Context) error {
	prodTeamID := paramFromContext(c).ProdTeamID
	res, err := a.App.GetProductionTeam(prodTeamID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (a API) ProductionTeamJoinHub(c echo.Context) error {
	hub := model.Hub{}
	err := c.Bind(&hub)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	prodTeamID := paramFromContext(c).ProdTeamID
	res, err := a.App.ProductionTeamJoinHub(prodTeamID, hub.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}