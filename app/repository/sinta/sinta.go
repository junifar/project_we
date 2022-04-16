package sinta

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
	sintarpm "project_we/app/repository/sinta/model"
)

type ISinta interface {
	GetPartnerSinta(ctx *context.Context, nidn string) (res sintarpm.SintaResponse, err error)
	GetSintaByPersonalID(ctx *context.Context, personalID int64) (res []model.Sinta, err error)
	InsertSinta(ctx *context.Context, payload model.Sinta) (err error)
	UpdateSinta(ctx *context.Context, payload model.Sinta) (err error)
}
