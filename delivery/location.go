package delivery

import (
	"project_we/app/model"
	rsp "project_we/app/pkg/response"
)

// City handler
func (impl *Deliveries) City() {
	ctx := impl.Ctx
	var city model.Cities

	errs := bindCity(ctx, &city)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	res, err := impl.Location.GetCity(ctx, city)
	if err != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	rsp.WriteResponse(&impl.Controller, nil, res)
}

// CityCreate handler
func (impl *Deliveries) CityCreate() {
	ctx := impl.Ctx
	var reqCity model.Cities

	errs := bindCity(ctx, &reqCity)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	errs = impl.Location.CreateCity(ctx, reqCity)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	rsp.WriteResponse(&impl.Controller, nil, nil)
}
