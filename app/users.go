package app

import "github.com/buivuanh/cubicasa/model"

func (a App) GetUser(id int) (*model.User, error){
	u,err := a.Store.User().Get(id)
	if err != nil{
		return nil, err
	}

	return u, nil
}

func (a App)  CreateUser(teamID int, u *model.User) (*model.User,error){
	res, err := a.Store.User().Create(teamID, u)
	if err != nil {
		return nil, err
	}

	return res, nil
}
