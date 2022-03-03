package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
	"time"
)

func (impl *UserUsecase) ListPersonel(ctx *context.Context, req model.PersonalsPayload) (res []model.Personals, errs errors.IError) {
	logs.Info("get personal list data")
	var limit, page int
	if req.Limit == 0 {
		limit = 10
	}
	res, err := impl.User.SelectPersonal(ctx, req.Personals, limit, page)
	if err != nil {
		logs.Error("failed get personal list data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}
	return
}

func (impl *UserUsecase) CreatePersonel(ctx *context.Context, req model.Personals) (errs errors.IError) {
	// TODO validate user input
	userID := int64(0)

	now := time.Now()

	req.CreateBy = userID
	req.CreateTime = now
	req.UpdateBy = userID
	req.UpdateTime = now

	_, err := impl.User.InsertPersonal(ctx, req)
	if err != nil {
		logs.Error("failed insert personal data :", err)
		return errors.New(constant.ErrorInternaly)
	}

	return nil
}

func (impl *UserUsecase) CheckAdmin(ctx *context.Context, req model.Personals) (res bool, errs errors.IError) {
	if req.UserTypeId == constant.PersonalTypeAdminID {
		res = true
	}
	return
}
