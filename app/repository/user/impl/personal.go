package impl

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

func (impl *UserRepository) SelectPersonal(ctx *context.Context, req model.Personals) (res []model.Personals, err error) {
	querySeter := impl.orm.QueryTable("personals")

	if req.ID != 0 {
		querySeter = querySeter.Filter("id", req.ID)
	}

	if req.Name != "" {
		querySeter = querySeter.Filter("name", req.Name)
	}

	if req.Username != "" {
		querySeter = querySeter.Filter("username", req.Username)
	}

	//if req.Password != "" {
	//	querySeter = querySeter.Filter("password", req.Password)
	//}

	_, err = querySeter.All(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (impl *UserRepository) InsertPersonal(ctx *context.Context, req model.Personals) (id int64, err error) {
	res := new(model.Personals)

	res.Name = req.Name
	res.Username = req.Username
	res.Password = req.Password
	res.CreateBy = req.CreateBy
	res.CreateTime = req.CreateTime
	res.UpdateBy = req.UpdateBy
	res.UpdateTime = req.UpdateTime

	id, err = impl.orm.Insert(res)
	if err != nil {
		return 0, err
	}

	return id, nil
}
