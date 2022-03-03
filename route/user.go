package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
)

func User(delivery *delivery.Deliveries) {
	Personal(delivery)
	Authorization(delivery)
}

func Personal(delivery *delivery.Deliveries) {
	//web.Router("/internal/v1/personal", delivery, "get:User")
	web.Router("/internal/v1/personal/create", delivery, "post:PersonalCreate")
}

func Authorization(delivery *delivery.Deliveries) {
	web.Router("/v1/login", delivery, "post:Login")
}
