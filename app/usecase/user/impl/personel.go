package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
	"time"
)

func (impl *UserUsecase) ListPersonel(ctx *context.Context, req model.PersonalsPayload) (res model.PersonalResults, errs errors.IError) {
	logs.Info("get personal list data")
	var limit, page int
	if req.Limit < 10 {
		limit = 10
	}
	page = req.Page - 1

	resData, err := impl.User.SelectPersonal(ctx, req.Personals, limit, page)
	if err != nil {
		logs.Error("failed get personal list data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	if len(resData) > limit {
		res.NextPage = req.Page + 1
		if req.Page-1 >= 0 {
			res.PreviousPage = req.Page - 1
		}
		res.Personals = resData[:len(resData)-1]
	} else {
		if req.Page-1 >= 0 {
			res.PreviousPage = req.Page - 1
		}
		res.Personals = resData
	}
	return
}

func (impl *UserUsecase) CreatePersonel(ctx *context.Context, req model.Personals) (errs errors.IError) {
	// TODO validate user input
	userID := ctx.Input.GetData(constant.ContextUserID).(int64)

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

func (impl *UserUsecase) CurrentUser(ctx *context.Context, req model.Personals) (res model.Personals, errs errors.IError) {
	userID := ctx.Input.GetData(constant.ContextUserID).(int64)

	personalPayload := model.Personals{ID: userID}
	personalList, err := impl.User.SelectPersonal(ctx, personalPayload, 0, 0)
	if err != nil || len(personalList) == 0 {
		logs.Error("failed get personal list data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	res = personalList[0]
	return
}

func (impl *UserUsecase) CheckAdmin(ctx *context.Context, req model.Sessions) (res bool, errs errors.IError) {
	for _, peran := range req.IDPeran {
		if peran == constant.PersonalTypeAdminID {
			res = true
			return
		}
	}
	return
}

func (impl *UserUsecase) CheckLecturer(ctx *context.Context, req model.Sessions) (res bool, errs errors.IError) {
	for _, peran := range req.IDPeran {
		if peran == constant.PersonalTypeLecturerID {
			res = true
			return
		}
	}
	return
}
