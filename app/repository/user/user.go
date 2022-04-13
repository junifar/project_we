package user

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

// IUser is user repository interface
type IUser interface {
	IPersonal
	ISession
	IdentitasPengguna
	IPeran
}

// IPersonal is personal repository interface
type IPersonal interface {
	SelectPersonal(ctx *context.Context, req model.Personals, limit, page int) (res []model.Personals, err error)
	InsertPersonal(ctx *context.Context, req model.Personals) (id int64, err error)
	UpdatePersonal(ctx *context.Context, req model.Personal) (res model.Personal, err error)

	SelectPersonalByIDPersonal(ctx *context.Context, idPersonal int64) (res model.Personal, err error)
	SelectPersonalByFilter(ctx *context.Context, filter model.PersonalFilter, limit, page int) (res []model.Personal, err error)
}

type IdentitasPengguna interface {
	SelectIdentitasPenggunaByUserName(ctx *context.Context, username string) (res []model.IdentitasPengguna, err error)
	InsertIdentitasPengguna(ctx *context.Context, req model.Personals) (id int64, err error)
}

type IPeran interface {
	SelectPeranByIDPersonal(ctx *context.Context, idPersonal int64) (res []model.Peran, err error)
}

// ISession is session repository interface
type ISession interface {
	SetSession(ctx *context.Context, req model.Sessions, uuid string) (string, error)
	GetSession(ctx *context.Context, cookie string) (model.Sessions, error)
	RemoveSession(ctx *context.Context, cookie string) error
}
