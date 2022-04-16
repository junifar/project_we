package sinta

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/repository/sinta/model"
)

type ISinta interface {
	GetPartnerSinta(ctx *context.Context, nidn string) (res model.SintaResponse, err error)
}
