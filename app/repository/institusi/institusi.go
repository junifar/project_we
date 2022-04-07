package institusi

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

type IInstitusi interface {
	GetInstitusiByInstitusiID(ctx *context.Context, institusiID int64) (res model.Institusi, err error)
}
