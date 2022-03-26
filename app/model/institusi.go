package model

import "time"

type (
	Institusi struct {
		IDInstitusi      int64     `orm:"column:id_institusi"`
		KdJenisInstitusi string    `orm:"column:kd_jenis_institusi"`
		NamaInstitusi    string    `orm:"column:nama_institusi"`
		Alamat           string    `orm:"column:alamat"`
		KdKota           string    `orm:"column:kd_kota"`
		KdPos            string    `orm:"column:kd_pos"`
		Telepon          string    `orm:"column:telepon"`
		Fax              string    `orm:"column:fax"`
		Surel            string    `orm:"column:surel"`
		Website          string    `orm:"column:website"`
		IDPdpt           int64     `orm:"column:id_pdpt"`
		TglCreated       time.Time `orm:"column:tgl_created"`
		TglUpdated       time.Time `orm:"column:tgl_updated"`
		KdStsAktif       string    `orm:"column:kd_sts_aktif"`
		Level            int16     `orm:"column:level"`
		IDInstitusiInduk int64     `orm:"column:id_institusi_induk"`
		TokenVerifikasi  string    `orm:"column:token_verifikasi"`
	}
)
