package delivery

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"

	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
)

func bindCity(ctx *context.Context, reqCity *model.Cities) errors.IError {
	err := ctx.Input.Bind(&reqCity.Name, "name")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	return nil
}

func bindPersonal(ctx *context.Context, req *model.Personals) (errs errors.IError) {
	err := ctx.Input.Bind(&req.Username, "username")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Name, "name")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Password, "password")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	return nil
}

func bindPersonalUpdate(ctx *context.Context, req *model.Personal) (errs errors.IError) {
	err := ctx.Input.Bind(&req.NomorKtp, "nomor_ktp")
	if err != nil {
		logs.Error("failed binding input nomor_ktp", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Alamat, "alamat")
	if err != nil {
		logs.Error("failed binding input alamat", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.TempatLahir, "tempat_lahir")
	if err != nil {
		logs.Error("failed binding input tempat_lahir", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.NomorHp, "nomor_hp")
	if err != nil {
		logs.Error("failed binding input nomor_hp", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Surel, "surel")
	if err != nil {
		logs.Error("failed binding input surel", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	// optional form-data
	nomorTelepon := ctx.Input.Query("nomor_telepon")
	if nomorTelepon != "" {
		req.NomorTelepon = nomorTelepon
	}

	websitePersonal := ctx.Input.Query("website_personal")
	if websitePersonal != "" {
		req.WebsitePersonal = websitePersonal
	}

	return nil
}

func bindPersonalList(ctx *context.Context, req *model.PersonalsPayload) (errs errors.IError) {
	err := ctx.Input.Bind(&req.Personals.Username, "username")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Personals.Name, "name")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Personals.Password, "password")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Personals.Username, "username")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Limit, "limit")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Page, "page")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	return nil
}
