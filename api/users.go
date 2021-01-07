package api

import (
	"github.com/buivuanh/cubicasa/model"
	"github.com/labstack/echo"
	"net/http"
)

func (a *API) InitUser() {
	a.Routes.UserForTeam.POST("",a.CreateUser)
	a.Routes.User.GET("",a.GetUser)
}

func (a API) CreateUser(c echo.Context) error {
	teamID := paramFromContext(c).ProdTeamID
	u := model.User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := a.App.CreateUser(teamID, &u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (a API) GetUser(c echo.Context) error {
	userID := paramFromContext(c).UserID
	res, err := a.App.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
