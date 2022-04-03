package user

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
	"project_we/app/pkg/errors"
	userucm "project_we/app/usecase/user/model"
)

// IUser is user usecase interface
type IUser interface {
	IPersonel
	ISession
	IAuthorization
}

type IPersonel interface {
	ListPersonel(ctx *context.Context) (res []model.Personal, errs errors.IError)
	CurrentUser(ctx *context.Context) (res userucm.UserResponse, errs errors.IError)
	CreatePersonel(ctx *context.Context, req model.Personals) (errs errors.IError)
	CheckAdmin(ctx *context.Context, req model.Sessions) (res bool, errs errors.IError)
	CheckLecturer(ctx *context.Context, req model.Sessions) (res bool, errs errors.IError)
}

type ISession interface {
	CheckSession(ctx *context.Context, cookie string) (res model.Sessions, err errors.IError)
}

type IAuthorization interface {
	Login(ctx *context.Context, req model.Personals, uuid string) (cookie string, errs errors.IError)
}
