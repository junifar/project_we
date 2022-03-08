package impl

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
	"strconv"
	"strings"
	"time"
)

// SetSession is function to set session into cache
func (impl *UserRepository) SetSession(ctx *context.Context, req model.Sessions, uuid string) (string, error) {
	userIDString := strconv.FormatInt(req.IDPersonal, 10)
	userToken := hex.EncodeToString([]byte(userIDString))
	cookie := uuid + "-" + userToken

	key := fmt.Sprintf("session:%s", userToken)

	value, err := json.Marshal(req)
	if err != nil {
		return cookie, err
	}

	err = impl.Cache.Put(key, value, 8*time.Hour)
	if err != nil {
		return cookie, err
	}

	return cookie, nil
}

// GetSession is function to get session from cache
func (impl *UserRepository) GetSession(ctx *context.Context, cookie string) (model.Sessions, error) {
	res := model.Sessions{}
	splitToken := strings.Split(cookie, "-")
	userToken := splitToken[len(splitToken)-1]

	key := fmt.Sprintf("session:%s", userToken)

	isExist := impl.Cache.IsExist(key)
	if !isExist {
		return res, errors.New("cache does not exist")
	}

	cache := impl.Cache.Get(key).([]byte)

	err := json.Unmarshal(cache, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// RemoveSession is function to remove session from cache
func (impl *UserRepository) RemoveSession(ctx *context.Context, cookie string) error {
	splitToken := strings.Split(cookie, "-")
	userToken := splitToken[len(splitToken)-1]

	key := fmt.Sprintf("session:%s", userToken)

	isExist := impl.Cache.IsExist(key)
	if !isExist {
		return errors.New("cache does not exist")
	}

	err := impl.Cache.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
