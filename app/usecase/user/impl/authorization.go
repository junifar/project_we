package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
)

// Login is usecase to login
func (impl *UserUsecase) Login(ctx *context.Context, req model.Personals, uuid string) (cookie string, errs errors.IError) {
	logs.Info("get personal data by username")
	identitasPengguna, err := impl.User.SelectIdentitasPenggunaByUserName(ctx, req.Username)
	if err != nil || len(identitasPengguna) == 0 {
		logs.Error("failed get identitas pengguna data by username : %s with error : %+v", req.Username, err)
		errs = errors.New(constant.ErrorUserDoesntExist)
		return
	}

	userData := identitasPengguna[0]

	//validate password
	if userData.Pswd != req.Password {
		logs.Error("password is incorrect")
		errs = errors.New(constant.ErrorIncorrectPassword)
		return
	}

	//save to session
	payloadSession := model.Sessions{
		UserID:   userData.IdPersonal,
		Name:     userData.NamaUser,
		Username: userData.NamaUser,
	}

	logs.Info("set session to cache for user id", userData.IdPersonal)
	cookie, err = impl.User.SetSession(ctx, payloadSession, uuid)
	if err != nil {
		logs.Error("failed set session to cache :", err)
		errs = errors.New(constant.ErrorInternaly)
		return
	}

	return
}
