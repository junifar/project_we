package user

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

// IUser is user repository interface
type IUser interface {
	IPersonal
	ISession
	IdentitasPengguna
}

// IPersonal is personal repository interface
type IPersonal interface {
	SelectPersonal(ctx *context.Context, req model.Personals, limit, page int) (res []model.Personals, err error)
	InsertPersonal(ctx *context.Context, req model.Personals) (id int64, err error)
}

type IdentitasPengguna interface {
	SelectIdentitasPenggunaByUserName(ctx *context.Context, username string) (res []model.IdentitasPengguna, err error)
	InsertIdentitasPengguna(ctx *context.Context, req model.Personals) (id int64, err error)
}

// ISession is session repository interface
type ISession interface {
	SetSession(ctx *context.Context, req model.Sessions, uuid string) (string, error)
	GetSession(ctx *context.Context, cookie string) (model.Sessions, error)
	RemoveSession(ctx *context.Context, cookie string) error
}
