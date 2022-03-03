package function

import (
	"fmt"

	"github.com/beego/beego/v2/adapter/cache"
	_ "github.com/beego/beego/v2/adapter/cache/redis"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

// InitCache is function to initiation cache
func InitCache() (cache.Cache, error) {
	cacheConnectionString, err := config.String("redis::cacheConnectionString")
	if err != nil {
		logs.Error("failed get config :", err)
		return nil, err
	}

	cache, err := cache.NewCache("redis", fmt.Sprintf(`{"conn":"%s"}`, cacheConnectionString))
	if err != nil {
		logs.Error("failed cache registering driver :", err)
		return nil, err
	}

	return cache, nil
}
