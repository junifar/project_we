package impl

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

const (
	BaseQueryGetInstitusi = `
								select id_institusi,
									   kd_jenis_institusi,
									   nama_institusi,
									   alamat,
									   kd_kota,
									   kd_pos,
									   telepon,
									   fax,
									   surel,
									   website,
									   id_pdpt,
									   tgl_created,
									   tgl_updated,
									   kd_sts_aktif,
									   level,
									   id_institusi_induk,
									   token_verifikasi
								from tkt.institusi
							`
)

func (impl *institusiRepository) GetInstitusiByInstitusiID(ctx *context.Context, institusiID int64) (res model.Institusi, err error) {
	query := BaseQueryGetInstitusi + ` WHERE id_institusi = ?`

	err = impl.orm.Raw(query, institusiID).QueryRow(&res)
	if err == nil {
		return
	}
	return
}
