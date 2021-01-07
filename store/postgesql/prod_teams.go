package postgesql

import (
	"github.com/buivuanh/cubicasa/model"
	"time"
	"upper.io/db.v3/lib/sqlbuilder"
)

const prodTeamsTable string = "production_teams"

type ProdTeamStore struct {
	db sqlbuilder.Database
}

func NewProdTeamStore(db sqlbuilder.Database) ProdTeamStore {
	return ProdTeamStore{
		db: db,
	}
}

func (s ProdTeamStore) Get(id int) (*model.ProductionTeam, error) {
	pt := model.ProductionTeam{}
	err := s.db.SelectFrom(prodTeamsTable).Where("id", id).One(&pt)
	if err != nil {
		return nil, err
	}

	return &pt, nil
}

func (s ProdTeamStore) Create(pt *model.ProductionTeam) (*model.ProductionTeam, error) {
	if err := pt.IsValid(); err != nil{
		return nil, err
	}

	i := s.db.InsertInto(prodTeamsTable).Values(pt).Returning("id", "created_at").Iterator()
	defer i.Close()

	err := i.NextScan(&pt.ID, &pt.CreatedAt)
	if err != nil {
		return nil, err
	}

	return pt, nil
}

func (s ProdTeamStore) Update(pt *model.ProductionTeam) (*model.ProductionTeam, error) {
	if err := pt.IsValid();err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	pt.UpdatedAt = now

	_, err :=s.db.Update(prodTeamsTable).Set(*pt).Where("id", pt.ID).Exec()
	if err != nil {
		return nil, err
	}

	return pt, nil
}