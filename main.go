package main

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"project_we/delivery"
	"project_we/function"
	"project_we/route"

	locationrpi "project_we/app/repository/location/impl"

	locationuci "project_we/app/usecase/location/impl"
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

	logs.Info("initialize usecase")
	locationUC := locationuci.New(locationRP)
	logs.Info("initialize usecase location")

	// initialize delivery
	delivery := delivery.New(locationUC)
	logs.Info("initialize delivery")

	route.Common(delivery)
	route.Location(delivery)

	web.Run()
}
