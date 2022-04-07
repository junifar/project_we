package impl

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

const (
	BaseQueryGetProgramStudi = `select * from pdpt.program_studi`
)

func (impl *programStudiRepository) GetProgramStudiByProgramStudiID(ctx *context.Context, programStudiID int64) (res model.ProgramStudi, err error) {
	query := BaseQueryGetProgramStudi + ` WHERE id_program_studi = ?`

	err = impl.orm.Raw(query, programStudiID).QueryRow(&res)
	if err == nil {
		return
	}
	return
}
