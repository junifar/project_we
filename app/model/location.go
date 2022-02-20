package model

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// Cities model
type Cities struct {
	ID         int64 `orm:"auto;column(id)"`
	Name       string
	CreateBy   int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateBy   int64
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Cities))
}
