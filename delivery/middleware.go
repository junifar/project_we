package delivery

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/pkg/errors"
)

func (impl *Deliveries) CheckSession(ctx *context.Context) errors.IError {
	cookie, valid := ctx.GetSecureCookie(constant.CookieSecret, constant.CookieName)
	if !valid {
		return errors.New(constant.ErrorUnauthorized)
	}

	personalData, err := impl.User.CheckSession(ctx, cookie)
	if err != nil {
		return errors.New(constant.ErrorUnauthorized)
	}

	ctx.Input.SetData(constant.ContextUserID, personalData.ID)

	return nil
}

func (impl *Deliveries) CheckAdmin(ctx *context.Context) errors.IError {
	cookie, valid := ctx.GetSecureCookie(constant.CookieSecret, constant.CookieName)
	if !valid {
		return errors.New(constant.ErrorUnauthorized)
	}

	personalData, err := impl.User.CheckSession(ctx, cookie)
	if err != nil {
		return errors.New(constant.ErrorUnauthorized)
	}

	isAdmin, err := impl.User.CheckAdmin(ctx, personalData)
	if err != nil || !isAdmin {
		return errors.New(constant.ErrorNoPermission)
	}

	ctx.Input.SetData(constant.ContextUserID, personalData.ID)
	return nil
}
