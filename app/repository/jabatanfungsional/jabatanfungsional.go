package jabatanfungsional

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

type IJabatanFungsional interface {
	GetJabatanFungsionalByIDJabatanFungsional(ctx *context.Context, IDJabatanFungsional int) (res model.JabatanFungsional, err error)
}
