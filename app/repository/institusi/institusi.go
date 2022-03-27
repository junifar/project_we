package institusi

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

type IInstitusi interface {
	GetInstitusiByInstitusiID(ctx *context.Context, institusiID int64) (res model.Institusi, err error)
}
