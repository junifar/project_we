package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
)

// CheckSession is usecase to check user login session
func (impl *UserUsecase) CheckSession(ctx *context.Context, cookie string) (res model.Sessions, errs errors.IError) {
	logs.Info("get session from cache for session cookie", cookie)
	res, err := impl.User.GetSession(ctx, cookie)
	if err != nil {
		logs.Error("failed get session from cache :", err)
		errs = errors.New(constant.ErrorDataNotFoundCache)
		return
	}
	return
}
