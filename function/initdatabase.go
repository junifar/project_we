package function

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/lib/pq"
)

// InitDB is function to initiation database
func InitDB() error {
	dbAlias := "default"
	//force := false
	//verbose := true
	driverName := "postgres"

	err := orm.RegisterDriver(driverName, orm.DRPostgres)
	if err != nil {
		logs.Error("failed orm registering driver :", err)
		return err
	}

	dbConnectionString, err := config.String("database::dbConnectionString")
	if err != nil {
		logs.Error("failed get config :", err)
		return err
	}

	err = orm.RegisterDataBase(dbAlias, driverName, dbConnectionString)
	if err != nil {
		logs.Error("failed orm registering database :", err)
		return err
	}

	//err = orm.RunSyncdb(dbAlias, force, verbose)
	//if err != nil {
	//	logs.Error("failed running syncronization database :", err)
	//	return err
	//}

	return nil
}
