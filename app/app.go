package app

import (
	"github.com/buivuanh/cubicasa/model"
	"github.com/buivuanh/cubicasa/store"
	"github.com/buivuanh/cubicasa/store/postgesql"
)

type App struct {
	Store  store.Store

}

func New(conf model.AppSettings) (*App, error) {
	sqlStore, err := postgesql.NewStore(conf.DatabaseSettings)
	if err != nil {
		return nil, err
	}

	a := App{
		Store: sqlStore,
	}

	return &a, nil
}