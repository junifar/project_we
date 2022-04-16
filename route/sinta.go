package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
)

func Sinta(delivery *delivery.Deliveries) {
	web.Router("/v1/sinta", delivery, "get:SyncSinta")
}
