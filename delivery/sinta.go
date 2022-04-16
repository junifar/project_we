package delivery

import rsp "project_we/app/pkg/response"

func (impl *Deliveries) SyncSinta() {
	ctx := impl.Ctx
	res, errs := impl.Sinta.SyncSinta(ctx)
	if errs != nil {
		rsp.WriteResponse(&impl.Controller, errs, nil)
		return
	}

	rsp.WriteResponse(&impl.Controller, nil, res)
}
