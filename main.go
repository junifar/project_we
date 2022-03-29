package main

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"os"

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

	logs.Info("initialize cache")
	cache, err := function.InitCache()
	if err != nil {
		return
	}

	logs.Info("initialize repository")
	rpList := repository.Init(repository.Dependencies{
		ORM:   orms,
		Cache: cache,
	})

	logs.Info("initialize usecase")
	ucList := usecase.Init(usecase.Repository{
		LocationRP:  rpList.LocationRP,
		UserRP:      rpList.UserRP,
		DosenRP:     rpList.DosenRP,
		InstitusiRP: rpList.InstitusiRP,
	})

	// initialize delivery
	deliveries := delivery.New(delivery.Usecase{
		LocationUC: ucList.LocationUC,
		UserUC:     ucList.UserUC,
	})
	logs.Info("initialize delivery")

	route.Init(deliveries)

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
	}))

	//web.Run()
	logs.Info("Starting server at port %s\n", os.Getenv("PORT"))
	web.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
