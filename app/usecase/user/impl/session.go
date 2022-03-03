package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
)

// CheckSession is usecase to check user login session
func (impl *UserUsecase) CheckSession(ctx *context.Context, cookie string) (res model.Personals, errs errors.IError) {
	logs.Info("get session from cache for session cookie", cookie)
	sessionData, err := impl.User.GetSession(ctx, cookie)
	if err != nil {
		logs.Error("failed get session from cache :", err)
		errs = errors.New(constant.ErrorDataNotFoundCache)
		return
	}

	logs.Info("get user data by user id", sessionData.UserID)
	personalPayloads := model.Personals{ID: sessionData.UserID}
	personalList, err := impl.User.SelectPersonal(ctx, personalPayloads, 0, 0)
	if err != nil {
		logs.Error("failed get user data :", err)
		errs = errors.New(constant.ErrorDataNotFoundDB)
		return
	}
	if len(personalList) == 1 {
		res = personalList[0]
	}
	return
}
