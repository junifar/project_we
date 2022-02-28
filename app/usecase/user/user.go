package user

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
	"project_we/app/pkg/errors"
)

// IUser is user usecase interface
type IUser interface {
	IPersonel
	ISession
}

type IPersonel interface {
	CreatePersonel(ctx *context.Context, req model.Personals) (err errors.IError)
	CheckAdmin(ctx *context.Context, req model.Personals) (res bool, errs errors.IError)
}

type ISession interface {
	CheckSession(ctx *context.Context, cookie string) (res model.Personals, err errors.IError)
}
