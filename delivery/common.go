package delivery

import (
	rsp "project_we/app/pkg/response"
)

// Index handler
func (impl *Deliveries) Index() {
	ctx := impl.Ctx

	ctx.WriteString("OK")
}

// Ping handler
func (impl *Deliveries) Ping() {
	rsp.WriteResponse(&impl.Controller, nil, "pong")
}
