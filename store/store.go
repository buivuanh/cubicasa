package store

import "github.com/buivuanh/cubicasa/model"

type Store interface {
	User() User
	ProductionTeam() ProductionTeam
	Hub() Hub
}

type User interface {
	Get(id int) (*model.User, error)
	Create(teamID int, u *model.User) (*model.User, error)
}

type ProductionTeam interface {
	Get(id int) (*model.ProductionTeam, error)
	Create(pt *model.ProductionTeam) (*model.ProductionTeam, error)
	Update(pt *model.ProductionTeam) (*model.ProductionTeam, error)
}

type Hub interface {
	Get(id int) (*model.Hub, error)
	Create(h *model.Hub) (*model.Hub, error)
}
