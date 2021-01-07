package app

import "github.com/buivuanh/cubicasa/model"

func (a App) GetHub(id int) (*model.Hub, error) {
	h, err := a.Store.Hub().Get(id)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (a App) CreateHub(h *model.Hub) (*model.Hub, error) {
	res, err := a.Store.Hub().Create(h)
	if err != nil {
		return nil, err
	}

	return res, nil
}
