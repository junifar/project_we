package impl

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

const (
	BaseQueryGetJabatanFungsional = `
										select id_jabatan_fungsional,
											   jabatan_fungsional,
											   tgl_created,
											   tgl_updated,
											   kd_sts_aktif
										from pdpt.jabatan_fungsional
									`
)

func (impl *jabatanFungsionalRepository) GetJabatanFungsionalByIDJabatanFungsional(ctx *context.Context, IDJabatanFungsional int) (res model.JabatanFungsional, err error) {
	query := BaseQueryGetJabatanFungsional + ` WHERE id_jabatan_fungsional = ?`

	err = impl.orm.Raw(query, IDJabatanFungsional).QueryRow(&res)
	if err == nil {
		return
	}
	return
}
