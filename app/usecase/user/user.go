package user

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
	"project_we/app/pkg/errors"
)

// IUser is user usecase interface
type IUser interface {
	Personel
}

type Personel interface {
	CreatePersonel(ctx *context.Context, req model.Personals) (err errors.IError)
}
