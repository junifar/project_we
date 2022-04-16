package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
)

func Init(delivery *delivery.Deliveries) {
	web.InsertFilter("*", web.BeforeRouter, delivery.InitContext)
	Common(delivery)
	Location(delivery)
	User(delivery)
	Sinta(delivery)
}
