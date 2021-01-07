package api

import (
	"github.com/buivuanh/cubicasa/app"
	"github.com/labstack/echo"
)

type Routes struct {
	Root *echo.Echo

	User        *echo.Group // "api/users/:userID"
	UserForTeam *echo.Group // "api/prodteams/:prodteamID/users"

	ProdTeam  *echo.Group // "api/prodteams/:prodteamID"
	ProdTeams *echo.Group // "api/prodteams"

	Hub  *echo.Group // "api/hubs/:hubID"
	Hubs *echo.Group // "api/hubs"
}

type API struct {
	App    *app.App
	Routes Routes
}

func New(ap *app.App, root *echo.Echo) (*API, error) {
	a := &API{
		App:    ap,
		Routes: Routes{},
	}

	a.Routes.Root = root
	a.Routes.ProdTeam = a.Routes.Root.Group("/prodteams/:prodTeamID")
	a.Routes.ProdTeams = a.Routes.Root.Group("/prodteams")

	a.Routes.UserForTeam = a.Routes.ProdTeam.Group("/users")
	a.Routes.User = a.Routes.Root.Group("/users/:userID")

	a.Routes.Hub = a.Routes.Root.Group("/hubs/:hubID")
	a.Routes.Hubs = a.Routes.Root.Group("/hubs")

	a.InitHub()
	a.InitProductionTeam()
	a.InitUser()

	return a, nil
}
