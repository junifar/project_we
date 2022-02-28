package impl

import (
	"github.com/beego/beego/v2/adapter/cache"
	"github.com/beego/beego/v2/adapter/orm"
	userrp "project_we/app/repository/user"
)

//UserRepository -
type UserRepository struct {
	orm   orm.Ormer
	Cache cache.Cache
}

// New user contructor
func New(orm orm.Ormer, Cache cache.Cache) userrp.IUser {
	return &UserRepository{
		orm:   orm,
		Cache: Cache,
	}
}
