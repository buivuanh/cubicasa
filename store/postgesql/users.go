package postgesql

import (
	"github.com/buivuanh/cubicasa/model"
	"upper.io/db.v3/lib/sqlbuilder"
)

const usersTable string = "users"

type UserStore struct {
	db sqlbuilder.Database
}

func NewUserStore(db sqlbuilder.Database) UserStore {
	return UserStore{
		db: db,
	}
}

func (s UserStore) Get(id int) (*model.User, error) {
	u := model.User{}
	err := s.db.SelectFrom(usersTable).Where("id", id).One(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s UserStore) Create(teamID int, u *model.User) (*model.User, error) {
	u.TeamID = teamID
	if err := u.IsValid(); err != nil{
		return nil, err
	}

	i := s.db.InsertInto(usersTable).Values(u).Returning("id", "created_at").Iterator()
	defer i.Close()

	err := i.NextScan(&u.UserID, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return u, nil
}