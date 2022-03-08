package impl

import (
	"github.com/beego/beego/v2/adapter/context"
	"project_we/app/model"
)

func (impl *UserRepository) SelectIdentitasPengguna(ctx *context.Context, req model.IdentitasPengguna, limit, page int) (res []model.IdentitasPengguna, err error) {
	querySetter := impl.orm.QueryTable("pengguna.identitas_pengguna")

	if req.IdPersonal != 0 {
		querySetter = querySetter.Filter("id_personal", req.IdPersonal)
	}

	if req.NamaUser != "" {
		querySetter = querySetter.Filter("nama_user", req.NamaUser)
	}

	_, err = querySetter.All(&res)
	if err != nil {
		return nil, err
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
