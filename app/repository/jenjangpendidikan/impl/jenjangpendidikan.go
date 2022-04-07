package impl

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

const (
	BaseQueryGetJenjangPendidikan = `select * from pdpt.jenjang_pendidikan`
)

func (impl *jenjangPendidikanRepository) GetJenjangPendidikanByJenjangPendidikanID(ctx *context.Context, jenjangPendidikanID int64) (res model.JenjangPendidikan, err error) {
	query := BaseQueryGetJenjangPendidikan + ` WHERE id_jenjang_pendidikan = ?`

	err = impl.orm.Raw(query, jenjangPendidikanID).QueryRow(&res)
	if err == nil {
		return
	}
	return
}
