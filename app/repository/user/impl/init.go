package impl

import (
	"github.com/beego/beego/v2/adapter/orm"
	userrp "project_we/app/repository/user"
)

//UserRepository -
type UserRepository struct {
	orm orm.Ormer
}

// New user contructor
func New(orm orm.Ormer) userrp.IUser {
	return &UserRepository{
		orm: orm,
	}
}
