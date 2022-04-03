package impl

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

const (
	BaseQueryGetPersonal = `
								select 
									   id_personal,
									   id_personal_uuid,
									   nama,
									   nomor_ktp,
									   alamat,
									   tempat_lahir,
									   tanggal_lahir,
									   nomor_telepon,
									   nomor_hp,
									   surel,
									   website_personal,
									   id_institusi,
									   tgl_updated,
									   tgl_created
								from tkt.personal
							`
)

func (impl *UserRepository) SelectPersonal(ctx *context.Context, req model.Personals, limit, page int) (res []model.Personals, err error) {
	querySeter := impl.orm.QueryTable("personals")

	if req.ID != 0 {
		querySeter = querySeter.Filter("id", req.ID)
	}

	if req.Name != "" {
		querySeter = querySeter.Filter("name", req.Name)
	}

	if req.Username != "" {
		querySeter = querySeter.Filter("username", req.Username)
	}

	_, err = querySeter.Limit(limit + 1).Offset(page * limit).All(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (impl *UserRepository) InsertPersonal(ctx *context.Context, req model.Personals) (id int64, err error) {
	res := new(model.Personals)

	res.Name = req.Name
	res.Username = req.Username
	res.Password = req.Password
	res.CreateBy = req.CreateBy
	res.CreateTime = req.CreateTime
	res.UpdateBy = req.UpdateBy
	res.UpdateTime = req.UpdateTime

	id, err = impl.orm.Insert(res)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (impl *UserRepository) SelectPersonalByIDPersonal(ctx *context.Context, idPersonal int64) (res model.Personal, err error) {
	query := BaseQueryGetPersonal + ` WHERE id_personal = ?`

	err = impl.orm.Raw(query, idPersonal).QueryRow(&res)
	if err == nil {
		return
	}
	return
}

func (impl *UserRepository) SelectPersonalByFilter(ctx *context.Context, filter model.PersonalFilter, limit, page int) (res []model.Personal, err error) {
	_, err = impl.orm.Raw(BaseQueryGetPersonal).QueryRows(&res)
	if err == nil {
		return
	}
	return
}
