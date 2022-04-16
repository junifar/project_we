package main

import (
	"fmt"
	"os"
	"project_we/builder"

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

	logs.Info("initialize cache")
	cache, err := function.InitCache()
	if err != nil {
		return
	}

	logs.Info("initialize repository")
	rpList := repository.Init(repository.Dependencies{
		ORM:   orms,
		Cache: cache,
		Curl:  builder.BuildHTTPRequestor(),
	})

	logs.Info("initialize usecase")
	ucList := usecase.Init(usecase.Repository{
		LocationRP:          rpList.LocationRP,
		UserRP:              rpList.UserRP,
		DosenRP:             rpList.DosenRP,
		InstitusiRP:         rpList.InstitusiRP,
		JenjangPendidikanRP: rpList.JenjangPendidikanRP,
		ProgramStudiRP:      rpList.ProgramStudiRP,
		JabatanFungsionalRP: rpList.JabatanFungsionalRP,
		SintaRP:             rpList.SintaRP,
	})

	// initialize delivery
	deliveries := delivery.New(delivery.Usecase{
		LocationUC: ucList.LocationUC,
		UserUC:     ucList.UserUC,
		SintaUC:    ucList.SintaUC,
	})
	logs.Info("initialize delivery")

	route.Init(deliveries)

	//web.Run()
	logs.Info("Starting server at port %s\n", os.Getenv("PORT"))
	web.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
