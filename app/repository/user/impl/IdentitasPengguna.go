package impl

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

const (
	BaseQueryGetIdentitasPengguna = `
										select 
											id_personal, 
											nama_user, 
											pswd, 
											tgl_data, 
											kd_sts_pengguna, 
											id_institusi, 
											tgl_updated, 
											tgl_created
										from pengguna.identitas_pengguna
									`
)

func (impl *UserRepository) SelectIdentitasPenggunaByUserName(ctx *context.Context, username string) (res []model.IdentitasPengguna, err error) {
	query := BaseQueryGetIdentitasPengguna + ` WHERE nama_user = ? `
	_, err = impl.orm.Raw(query, username).QueryRows(&res)
	if err == nil {
		return
	}
	return
}

func (impl *UserRepository) InsertIdentitasPengguna(ctx *context.Context, req model.Personals) (id int64, err error) {
	id, err = impl.orm.Insert(req)
	if err != nil {
		return 0, err
	}
	return
}
