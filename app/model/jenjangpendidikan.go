package model

import "time"

type JenjangPendidikan struct {
	IDJenjangPendidikan int64     `orm:"column:id_jenjang_pendidikan"`
	JenjangPendidikan   string    `orm:"column:jenjang_pendidikan"`
	Degree              string    `orm:"column:degree"`
	TglCreated          time.Time `orm:"column:tgl_created"`
	TglUpdated          time.Time `orm:"column:tgl_updated"`
	KdStsAktif          string    `orm:"column:kd_sts_aktif"`
	IDTingkatPendidikan int       `orm:"column:id_tingkat_pendidikan"`
}
