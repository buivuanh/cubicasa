package postgesql

import (
	"database/sql"
	"github.com/buivuanh/cubicasa/model"
	"github.com/buivuanh/cubicasa/store"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

type Store struct {
	db sqlbuilder.Database

	userStore     store.User
	prodTeamStore store.ProductionTeam
	hubStore      store.Hub
}

func NewStore(conf model.DatabaseSettings) (*Store, error) {
	dbx, err := sql.Open("postgres", conf.ConnectionString)
	if err != nil {
		return nil, err
	}

	db, err := postgresql.New(dbx)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(32)

	// try first query
	_, err = db.Exec("select 1")
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
		userStore: NewUserStore(db),
		prodTeamStore: NewProdTeamStore(db),
		hubStore: NewHubStore(db),
	}, nil
}

func (s Store) User() store.User {
	return s.userStore
}

func (s Store) ProductionTeam() store.ProductionTeam {
	return s.prodTeamStore
}

func (s Store) Hub() store.Hub {
	return s.hubStore
}
