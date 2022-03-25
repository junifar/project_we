package impl

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

const (
	BaseQueryGetDosen = `
							select nidn,
								   golongan,
								   pangkat,
								   kd_perguruan_tinggi,
								   id_program_studi,
								   id_fakultas,
								   id_jurusan,
								   kd_sts_aktif,
								   kd_jenjang_pendidikan_program_studi,
								   id_personal,
								   id_jenjang_pendidikan_tertinggi,
								   no_sertifikat_dosen,
								   id_jabatan_fungsional,
								   no_pegawai,
								   id_pdpt,
								   tgl_updated,
								   tgl_created
							from pdpt.dosen
						`
)

func (d *dosenRepository) SelectDosenByPersonalID(ctx *context.Context, personalID int64) (res model.Dosen, err error) {
	query := BaseQueryGetDosen + ` WHERE id_personal = ?`

	err = d.orm.Raw(query, personalID).QueryRow(&res)
	if err == nil {
		return
	}
	return

	return
}
