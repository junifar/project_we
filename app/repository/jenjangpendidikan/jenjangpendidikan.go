package jenjangpendidikan

import (
	"project_we/app/model"

	"github.com/beego/beego/v2/server/web/context"
)

type IJenjangPendidikan interface {
	GetJenjangPendidikanByJenjangPendidikanID(ctx *context.Context, jenjangPendidikanID int64) (res model.JenjangPendidikan, err error)
}
