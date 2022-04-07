package programstudi

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

type IProgramStudi interface {
	GetProgramStudiByProgramStudiID(ctx *context.Context, programStudiID int64) (res model.ProgramStudi, err error)
}
