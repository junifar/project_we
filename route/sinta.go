package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
)

func AccessSinta(delivery *delivery.Deliveries) {
	web.InsertFilter("/v1/sinta/sync", web.BeforeRouter, delivery.MustLecturer)
}

func Sinta(delivery *delivery.Deliveries) {
	AccessSinta(delivery)
	web.Router("/v1/sinta/sync", delivery, "post:SyncSinta")
}
