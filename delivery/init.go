package delivery

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/google/uuid"
	"project_we/app/constant"
	rsp "project_we/app/pkg/response"
	locationuc "project_we/app/usecase/location"
	useruc "project_we/app/usecase/user"
	"time"
)

// Deliveries is delivery dependencies
type Deliveries struct {
	web.Controller
	Location locationuc.ILocation
	User     useruc.IUser
}

// Usecase is usecase dependencies
type Usecase struct {
	LocationUC locationuc.ILocation
	UserUC     useruc.IUser
}

// New is delivery constructor
func New(usecase Usecase) *Deliveries {
	return &Deliveries{
		Location: usecase.LocationUC,
		User:     usecase.UserUC,
	}
}

func (impl *Deliveries) InitContext(ctx *context.Context) {
	ctx.Input.SetData(constant.ContextBirthTime, time.Now())
	ctx.Input.SetData(constant.ContextUUID, uuid.New().String())
}

func (impl *Deliveries) MustAdmin(ctx *context.Context) {
	errs := impl.CheckAdmin(ctx)
	if errs != nil {
		rsp.WriteResponseFilter(ctx, errs, nil)
		return
	}
}
