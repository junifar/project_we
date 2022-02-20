package location

import (
	"github.com/beego/beego/v2/server/web/context"

	"project_we/app/model"
	"project_we/app/pkg/errors"
)

// ILocation is location usecase interface
type ILocation interface {

	//location usecase
	GetCity(ctx *context.Context, req model.Cities) ([]model.Cities, errors.IError)
	CreateCity(ctx *context.Context, req model.Cities) errors.IError
}
