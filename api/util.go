package api

import (
	"github.com/labstack/echo"
	"strconv"
)

type PathParam struct {
	UserID     int
	ProdTeamID int
	HubID      int
}

func paramFromContext(c echo.Context) PathParam {
	pp := PathParam{}
	userID, err := strconv.Atoi(c.Param("userID"))
	if err == nil {
		pp.UserID = userID
	}

	prodTeamID, err := strconv.Atoi(c.Param("prodTeamID"))
	if err == nil {
		pp.ProdTeamID = prodTeamID
	}

	hubID, err := strconv.Atoi(c.Param("hubID"))
	if err == nil {
		pp.HubID = hubID
	}

	return pp
}