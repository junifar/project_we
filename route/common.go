package route

import (
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
	"project_we/function"
)

func Common(delivery *delivery.Deliveries) {
	web.Router("/", delivery, "get:Index")
	web.Router("/ping", delivery, "get:Ping")
	web.ErrorHandler("404", function.PageNotFound)
}
