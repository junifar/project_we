package model

import "time"

type (
	JabatanFungsional struct {
		IdJabatanFungsional int       `orm:"column:id_jabatan_fungsional"`
		JabatanFungsional   string    `orm:"column:jabatan_fungsional"`
		TglCreated          time.Time `orm:"column:tgl_created"`
		TglUpdated          time.Time `orm:"column:tgl_updated"`
		KdStsAktif          string    `orm:"column:kd_sts_aktif"`
	}
)
