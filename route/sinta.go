package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
)

func AccessSinta(delivery *delivery.Deliveries) {
	web.InsertFilter("/v1/dosen/sinta/sync", web.BeforeRouter, delivery.MustLecturer)
	web.InsertFilter("/v1/dosen/sinta", web.BeforeRouter, delivery.MustLecturer)
}

func Sinta(delivery *delivery.Deliveries) {
	AccessSinta(delivery)
	web.Router("/v1/dosen/sinta/sync", delivery, "post:SyncSinta")
	web.Router("/v1/dosen/sinta", delivery, "get:GetSinta")
}
