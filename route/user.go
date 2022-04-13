package route

import (
	"project_we/delivery"

	"github.com/beego/beego/v2/server/web"
)

func User(delivery *delivery.Deliveries) {
	AccessUser(delivery)
	Personal(delivery)
	Authorization(delivery)
}

func AccessUser(delivery *delivery.Deliveries) {
	web.InsertFilter("/internal/v1/users", web.BeforeRouter, delivery.MustAdmin)

	web.InsertFilter("/v1/dosen/identitas", web.BeforeRouter, delivery.MustLecturer)
	web.InsertFilter("/v1/dosen/identitas/update", web.BeforeRouter, delivery.MustLecturer)
}

func Personal(delivery *delivery.Deliveries) {
	web.Router("/v1/dosen/identitas", delivery, "get:CurrentUser")
	web.Router("/v1/dosen/identitas/update", delivery, "post:PersonalUpdate")

	web.Router("/internal/v1/users", delivery, "get:PersonalList")
	web.Router("/internal/v1/users/create", delivery, "post:PersonalCreate")
}

func Authorization(delivery *delivery.Deliveries) {
	web.Router("/v1/login", delivery, "post:Login")
	web.Router("/v1/logout", delivery, "get:Logout")

	web.Router("/internal/v1/login", delivery, "post:Login")
}
