package impl

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

const (
	QueryGetPeranByIDPersonal = `
							select distinct
								   pp.id_peran, 
								   pp.kd_aplikasi, 
								   pp.nama_peran, 
								   pp.keterangan, 
								   pp.kd_kelompok_peran, 
								   pp.tgl_created, 
								   pp.tgl_updated
							from pengguna.peran pp inner join pengguna.peran_pengguna ppp on pp.id_peran = ppp.id_peran
							where ppp.id_personal = ?
						`
)

func (impl *UserRepository) SelectPeranByIDPersonal(ctx *context.Context, idPersonal int64) (res []model.Peran, err error) {
	_, err = impl.orm.Raw(QueryGetPeranByIDPersonal, idPersonal).QueryRows(&res)
	if err == nil {
		return
	}
	return
}
