package location

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

// ILocation is location repository interface
type ILocation interface {
	//city repository
	InsertCity(ctx *context.Context, req model.Cities) (int64, error)
	SelectCity(ctx *context.Context, req model.Cities) ([]model.Cities, error)
}
