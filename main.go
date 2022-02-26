package main

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	"project_we/app/repository"
	"project_we/app/usecase"
	"project_we/delivery"
	"project_we/function"
	"project_we/route"
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
	rpList := repository.Init(repository.Dependencies{ORM: orms})

	logs.Info("initialize usecase")
	ucList := usecase.Init(usecase.Repository{
		LocationRP: rpList.LocationRP,
		UserRP:     rpList.UserRP,
	})

	// initialize delivery
	deliveries := delivery.New(delivery.Usecase{
		LocationUC: ucList.LocationUC,
		UserUC:     ucList.UserUC,
	})
	logs.Info("initialize delivery")

	route.Init(deliveries)

	web.Run()
}
