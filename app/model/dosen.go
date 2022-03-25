package model

import "time"

type (
	Dosen struct {
		NIDN                            string    `orm:"column(nidn)"`
		Golongan                        string    `orm:"column(golongan)"`
		Pangkat                         string    `orm:"column(pangkat)"`
		KdPerguruanTinggi               string    `orm:"column(kd_perguruan_tinggi)"`
		IDProgramStudi                  int64     `orm:"column(id_program_studi)"`
		IDFakultas                      int64     `orm:"column(id_fakultas)"`
		IDJurusan                       int64     `orm:"column(id_jurusan)"`
		KDSTSAktif                      string    `orm:"column(kd_sts_aktif)"`
		KDJenjangPendidikanProgramStudi string    `orm:"column(kd_jenjang_pendidikan_program_studi)"`
		IDPersonal                      int64     `orm:"column(id_personal)"`
		IDJenjangPendidikanTertinggi    int64     `orm:"column(id_jenjang_pendidikan_tertinggi)"`
		NoSertifikasiDosen              string    `orm:"column(no_sertifikasi_dosen)"`
		IDJabatanFungsional             int       `orm:"column(id_jabatan_fungsional)"`
		NoPegawai                       string    `orm:"column(no_pegawai)"`
		TglCreated                      time.Time `orm:"column(tgl_created)"`
		TglUpdated                      time.Time `orm:"column(tgl_updated)"`
		IDPDPT                          int64     `orm:"column(id_pdpt)"`
	}
)
