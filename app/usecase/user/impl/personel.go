package impl

import (
	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
	userucm "project_we/app/usecase/user/model"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

func (impl *UserUsecase) ListPersonel(ctx *context.Context) (res []model.Personal, errs errors.IError) {
	logs.Info("get personal list data")
	res, err := impl.User.SelectPersonalByFilter(ctx, model.PersonalFilter{}, 0, 0)
	if err != nil || len(res) == 0 {
		logs.Error("failed get personal list data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}
	return
}

func (impl *UserUsecase) CreatePersonel(ctx *context.Context, req model.Personals) (errs errors.IError) {
	// TODO validate user input
	userID := ctx.Input.GetData(constant.ContextUserID).(int64)

	now := time.Now()

	req.CreateBy = userID
	req.CreateTime = now
	req.UpdateBy = userID
	req.UpdateTime = now

	_, err := impl.User.InsertPersonal(ctx, req)
	if err != nil {
		logs.Error("failed insert personal data :", err)
		return errors.New(constant.ErrorInternaly)
	}

	return nil
}

func (impl *UserUsecase) GetPersonalByIDPersonal(ctx *context.Context, idPersonal int64) (model.Personal, errors.IError) {
	res, err := impl.User.SelectPersonalByIDPersonal(ctx, idPersonal)
	if err != nil {
		logs.Error("failed insert personal data :", err)
		return res, errors.New(constant.ErrorInternaly)
	}

	return res, nil
}

func (impl *UserUsecase) UpdatePersonal(ctx *context.Context, req model.Personal) (resUser userucm.UserResponse, err errors.IError) {
	req.IdPersonal = ctx.Input.GetData(constant.ContextUserID).(int64)

	_, errUpdate := impl.User.UpdatePersonal(ctx, req)
	if errUpdate != nil {
		logs.Error("failed update personal data :", err)
		return resUser, errors.New(constant.ErrorInternaly)
	}

	resUser, err = impl.CurrentUser(ctx)
	if err != nil {
		logs.Error("failed get personal data :", err)
		return resUser, errors.New(constant.ErrorInternaly)
	}

	return resUser, nil
}

func (impl *UserUsecase) CurrentUser(ctx *context.Context) (res userucm.UserResponse, errs errors.IError) {
	userID := ctx.Input.GetData(constant.ContextUserID).(int64)

	personalData, err := impl.User.SelectPersonalByIDPersonal(ctx, userID)
	if err != nil {
		logs.Error("failed get personal list data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	institusiData, err := impl.Institusi.GetInstitusiByInstitusiID(ctx, personalData.IdInstitusi)
	if err != nil {
		logs.Error("failed get institusi list data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	if personalData.IdPersonal > 0 {
		res.Nama = personalData.Nama
		res.Institusi = userucm.Institusi{
			IdInstitusi:   institusiData.IDInstitusi,
			NamaInstitusi: institusiData.NamaInstitusi,
		}
		res.Alamat = personalData.Alamat
		res.TempatLahir = personalData.TempatLahir
		res.TanggalLahir = personalData.TanggalLahir
		res.NomorKTP = personalData.NomorKtp
		res.NomorTelepon = personalData.NomorTelepon
		res.NomorHP = personalData.NomorHp
		res.Surel = personalData.Surel
		res.WebsitePersonal = personalData.WebsitePersonal
	}

	dosenData, err := impl.Dosen.SelectDosenByPersonalID(ctx, personalData.IdPersonal)
	if err != nil {
		logs.Error("failed get dosen data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	jenjangPendidikanData, err := impl.JenjangPendidikan.GetJenjangPendidikanByJenjangPendidikanID(ctx, dosenData.IDJenjangPendidikanTertinggi)
	if err != nil {
		logs.Error("failed get jenjang pendidikan data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	programStudiData, err := impl.ProgramStudi.GetProgramStudiByProgramStudiID(ctx, dosenData.IDProgramStudi)
	if err != nil {
		logs.Error("failed get program studi data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	jabatanFungsionalData, err := impl.JabatanFungsional.GetJabatanFungsionalByIDJabatanFungsional(ctx, dosenData.IDJabatanFungsional)
	if err != nil {
		logs.Error("failed get jabatan fungsional data :", err)
		return res, errors.New(constant.ErrorDataNotFoundDB)
	}

	if dosenData.NIDN != "" {
		res.NIDN = dosenData.NIDN

		res.ProgramStudi = userucm.ProgramStudi{
			IdProgramStudi:   dosenData.IDProgramStudi,
			NamaProgramStudi: programStudiData.NamaProgramStudi,
		}

		res.JenjangPendidikanTertinggi = userucm.JenjangPendidikanTertinggi{
			IdJenjangPendidikanTertinggi:   dosenData.IDJenjangPendidikanTertinggi,
			NamaJenjangPendidikanTertinggi: jenjangPendidikanData.JenjangPendidikan,
		}

		res.JabatanFungsional = userucm.JabatanFungsional{
			IdJabatanFungsional: jabatanFungsionalData.IdJabatanFungsional,
			JabatanFungsional:   jabatanFungsionalData.JabatanFungsional,
		}
	}

	return
}

func (impl *UserUsecase) CheckAdmin(ctx *context.Context, req model.Sessions) (res bool, errs errors.IError) {
	for _, peran := range req.IDPeran {
		if peran == constant.PersonalTypeAdminID {
			res = true
			return
		}
	}
	return
}

func (impl *UserUsecase) CheckLecturer(ctx *context.Context, req model.Sessions) (res bool, errs errors.IError) {
	for _, peran := range req.IDPeran {
		if peran == constant.PersonalTypeLecturerID {
			res = true
			return
		}
	}
	return
}
