package main

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
	"project_we/function"
	"project_we/route"

	locationrpi "project_we/app/repository/location/impl"
	userrpi "project_we/app/repository/user/impl"

	locationuci "project_we/app/usecase/location/impl"
	useruci "project_we/app/usecase/user/impl"
)

func init() {
	err := function.InitDB()
	if err != nil {
		return
	}
	logs.Info("initialize orm")
}

func main() {

	orms := orm.NewOrm()

	logs.Info("initialize repository")
	locationRP := locationrpi.New(orms)
	logs.Info("initialize repository location")
	userRP := userrpi.New(orms)
	logs.Info("initialize repository user")

	logs.Info("initialize usecase")
	locationUC := locationuci.New(locationRP)
	logs.Info("initialize usecase location")
	userUC := useruci.New(userRP)
	logs.Info("initialize usecase user")

	// initialize delivery
	delivery := delivery.New(locationUC, userUC)
	logs.Info("initialize delivery")

	route.Common(delivery)
	route.Location(delivery)
	route.User(delivery)

	web.Run()
}
