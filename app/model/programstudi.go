package model

import "time"

type ProgramStudi struct {
	IDProgramStudi      int64     `orm:"column:id_program_studi"`
	KdProgramStudi      string    `orm:"column:kd_program_studi"`
	KdPerguruanTinggi   string    `orm:"column:kd_perguruan_tinggi"`
	NamaProgramStudi    string    `orm:"column:nama_program_studi"`
	TglCreated          time.Time `orm:"column:tgl_created"`
	TglUpdated          time.Time `orm:"column:tgl_updated"`
	KdStsAktif          string    `orm:"column:kd_sts_aktif"`
	KdJenjangPendidikan int       `orm:"column:kd_jenjang_pendidikan"`
	IDProgramStudiPdpt  int64     `orm:"column:id_program_studi_pdpt"`
	KdProgramPendidikan int       `orm:"column:kd_program_pendidikan"`
}
