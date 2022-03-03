package delivery

import (
	"github.com/google/uuid"
	"project_we/app/constant"
	"project_we/app/model"
	delivery "project_we/app/pkg/cookie"
	rsp "project_we/app/pkg/response"
)

// PersonalCreate handler
func (impl *Deliveries) PersonalCreate() {
	ctx := impl.Ctx
	var req model.Personals

	errs := bindPersonal(ctx, &req)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	errs = impl.User.CreatePersonel(ctx, req)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	rsp.WriteResponse(&impl.Controller, nil, nil)
}

func (impl *Deliveries) PersonalList() {
	rsp.WriteResponse(&impl.Controller, nil, "test")
}

// Login handler
func (impl *Deliveries) Login() {
	ctx := impl.Ctx
	var req model.Personals

	errs := bindPersonal(ctx, &req)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	uniqueID := uuid.New().String()

	cookie, errs := impl.User.Login(ctx, req, uniqueID)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	delivery.SetSecureCookie(&impl.Controller, constant.CookieSecret, delivery.Cookie{
		Name:     constant.CookieName,
		Value:    cookie,
		MaxAge:   constant.CookieMaxAge,
		Secure:   false,
		HttpOnly: constant.CookieHttpOnly,
		SameSite: constant.CookieSameSite,
	})

	rsp.WriteResponse(&impl.Controller, nil, nil)
}

func (impl *Deliveries) Logout() {
	delivery.SetSecureCookie(&impl.Controller, constant.CookieSecret, delivery.Cookie{
		Name:     constant.CookieName,
		Value:    "",
		MaxAge:   -1 * constant.CookieMaxAge,
		Secure:   false,
		HttpOnly: constant.CookieHttpOnly,
		SameSite: constant.CookieSameSite,
	})
	rsp.WriteResponse(&impl.Controller, nil, nil)
}
