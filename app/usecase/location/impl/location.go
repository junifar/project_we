package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"time"

	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
)

// GetCity is usecase to get city
func (impl *LocationImpl) GetCity(ctx *context.Context, req model.Cities) ([]model.Cities, errors.IError) {
	// TODO validate user input
	var cityData []model.Cities

	logs.Info("select city")
	cityData, err := impl.Location.SelectCity(ctx, req)
	if err != nil {
		logs.Error("failed select city :", err)
		return cityData, errors.New(constant.ErrorDataNotFoundDB)
	}

	return cityData, nil
}

// CreateCity is usecase to create city
func (impl *LocationImpl) CreateCity(ctx *context.Context, req model.Cities) errors.IError {
	// TODO validate user input
	userID := int64(0)

	now := time.Now()

	req.CreateBy = userID
	req.CreateTime = now
	req.UpdateBy = userID
	req.UpdateTime = now

	_, err := impl.Location.InsertCity(ctx, req)
	if err != nil {
		logs.Error("failed insert product data :", err)
		return errors.New(constant.ErrorInternaly)
	}

	return nil
}
