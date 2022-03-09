package model

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type IdentitasPengguna struct {
	IdPersonal    int64     `orm:"column(id_personal)"`
	NamaUser      string    `orm:"column(nama_user)"`
	Pswd          string    `orm:"column(pswd)"`
	TglData       time.Time `orm:"column(tgl_data)"`
	KdStsPengguna string    `orm:"column(kd_sts_pengguna)"`
	IdInstitusi   int64     `orm:"column(id_institusi)"`
	TglUpdated    time.Time `orm:"column(tgl_updated)"`
	TglCreated    time.Time `orm:"column(tgl_created)"`
}

type Peran struct {
	IdPeran         int64
	KdApplikasi     string
	NamaPeran       string
	Keterangan      string
	KdKelompokPeran string
	TglCreated      time.Time
	TglUpdated      time.Time
}

type Personal struct {
	IdPersonal      int64
	NomorKtp        string
	Alamat          string
	TempatLahir     string
	TanggalLahir    time.Time
	NomorTelepon    string
	NomorHp         string
	Surel           string
	WebsitePersonal string
	TglCreated      time.Time
	TglUpdated      time.Time
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
