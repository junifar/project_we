package user

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

// IUser is user repository interface
type IUser interface {
	SelectPersonal(ctx *context.Context, req model.Personals) (res []model.Personals, err error)
	InsertPersonal(ctx *context.Context, req model.Personals) (id int64, err error)
}
