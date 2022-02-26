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
	err := ctx.Input.Bind(&req.Name, "usernsme")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Name, "name")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	err = ctx.Input.Bind(&req.Name, "password")
	if err != nil {
		logs.Error("failed binding input", err)
		return errors.New(constant.ErrorBindingParameter)
	}

	return nil
}
