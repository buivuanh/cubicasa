package postgesql

import (
	"github.com/buivuanh/cubicasa/model"
	"upper.io/db.v3/lib/sqlbuilder"
)

const hubsTable string = "hubs"

type HubStore struct {
	db sqlbuilder.Database
}

func NewHubStore(db sqlbuilder.Database) HubStore {
	return HubStore{
		db: db,
	}
}

func (s HubStore) Get(id int) (*model.Hub, error) {
	pt := model.Hub{}
	err := s.db.SelectFrom(hubsTable).Where("id", id).One(&pt)
	if err != nil {
		return nil, err
	}

	return &pt, nil
}

func (s HubStore) Create(h *model.Hub) (*model.Hub, error) {
	if err := h.IsValid(); err != nil{
		return nil, err
	}

	i := s.db.InsertInto(hubsTable).Values(h).Returning("id", "created_at").Iterator()
	defer i.Close()

	err := i.NextScan(&h.ID, &h.CreatedAt)
	if err != nil {
		return nil, err
	}

	return h, nil
}