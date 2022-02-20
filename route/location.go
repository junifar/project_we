package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
)

func Location(delivery *delivery.Deliveries) {
	City(delivery)
}

func City(delivery *delivery.Deliveries) {
	web.Router("/internal/v1/city", delivery, "get:City")
	web.Router("/internal/v1/city/create", delivery, "post:CityCreate")
}
