package delivery

import (
	"project_we/app/model"
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
