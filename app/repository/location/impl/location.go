package impl

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

// SelectCity is function to select city data
func (impl *LocationImpl) SelectCity(ctx *context.Context, req model.Cities) ([]model.Cities, error) {
	var res []model.Cities
	querySeter := impl.orm.QueryTable("cities")

	if req.ID != 0 {
		querySeter = querySeter.Filter("id", req.ID)
	}

	if req.Name != "" {
		querySeter = querySeter.Filter("name", req.Name)
	}

	_, err := querySeter.All(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// InsertCity is function to insert city data
func (impl *LocationImpl) InsertCity(ctx *context.Context, req model.Cities) (int64, error) {
	res := new(model.Cities)

	res.ID = req.ID
	res.Name = req.Name
	res.CreateBy = req.CreateBy
	res.CreateTime = req.CreateTime
	res.UpdateBy = req.UpdateBy
	res.UpdateTime = req.UpdateTime

	id, err := impl.orm.Insert(res)
	if err != nil {
		return 0, err
	}

	return id, nil
}
