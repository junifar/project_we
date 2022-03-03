package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
)

func User(delivery *delivery.Deliveries) {
	AccessUser(delivery)
	Personal(delivery)
	Authorization(delivery)
}

func AccessUser(delivery *delivery.Deliveries) {
	web.InsertFilter("/internal/v1/users", web.BeforeRouter, delivery.MustAdmin)
}

func Personal(delivery *delivery.Deliveries) {
	web.Router("/internal/v1/users", delivery, "get:PersonalList")
	web.Router("/internal/v1/users/create", delivery, "post:PersonalCreate")
}

func Authorization(delivery *delivery.Deliveries) {
	web.Router("/v1/login", delivery, "post:Login")
}
