package app

import "github.com/buivuanh/cubicasa/model"

func (a App) GetProductionTeam(id int) (*model.ProductionTeam, error) {
	pt, err := a.Store.ProductionTeam().Get(id)
	if err != nil {
		return nil, err
	}

	return pt, nil
}

func (a App) CreateProductionTeam(pt *model.ProductionTeam) (*model.ProductionTeam, error) {
	res, err := a.Store.ProductionTeam().Create(pt)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a App) ProductionTeamJoinHub(prodTeamID int, hubID int) (*model.ProductionTeam, error) {
	pt, err := a.GetProductionTeam(prodTeamID)
	if err != nil {
		return nil, err
	}

	pt.HubID = &hubID
	res, err := a.Store.ProductionTeam().Update(pt)
	if err != nil {
		return nil, err
	}

	return res, nil
}
