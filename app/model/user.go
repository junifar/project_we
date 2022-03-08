package model

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type IdentitasPengguna struct {
	ID            int64
	IdPersonal    string
	NamaUser      string
	Pswd          string
	TglData       time.Time
	KdStsPengguna string
	IdInstitusi   int64
	TglUpdated    time.Time
	TglCreated    time.Time
}

// Personals model
type Personals struct {
	ID         int64 `orm:"auto;column(id)"`
	Name       string
	Username   string `orm:"unique;index"`
	Password   string
	UserTypeId int64
	CreateBy   int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateBy   int64
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

type PersonalResults struct {
	Personals    []Personals `json:"list,omitempty"`
	NextPage     int         `json:"next_page,omitempty"`
	PreviousPage int         `json:"previous_page,omitempty""`
}

type PersonalsPayload struct {
	Personals Personals `json:"personals,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Page      int       `json:"page,omitempty"`
}

// UserTypes model
type UserTypes struct {
	ID         int64 `orm:"auto;column(id)"`
	Name       string
	CreateBy   int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateBy   int64
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Personals))
	orm.RegisterModel(new(UserTypes))
}
