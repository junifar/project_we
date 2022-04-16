package sinta

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/pkg/errors"
	sintaucm "project_we/app/usecase/sinta/model"
)

type ISinta interface {
	SyncSinta(ctx *context.Context) (res sintaucm.SintaResponse, errs errors.IError)
}
