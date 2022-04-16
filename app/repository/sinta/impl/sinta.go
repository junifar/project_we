package impl

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"project_we/app/repository/sinta/model"
)

func (impl *sintaRepository) GetPartnerSinta(ctx *context.Context, nidn string) (res model.SintaResponse, err error) {
	ctxt := context2.Background()

	apiKey, errs := config.String("sinta::apiKey")
	if errs != nil {
		logs.Error("failed get config apiKey :", errs)
		err = errs
		return
	}

	apiUrl, errs := config.String("sinta::apiURL")
	if errs != nil {
		logs.Error("failed get config apiURL :", errs)
		err = errs
		return
	}

	url := fmt.Sprintf("%s/select?nidn=%s&api_key=%s", apiUrl, nidn, apiKey)
	requestHttp := impl.curl.NewHttpRequest(http.MethodPost, url)

	logs.Info(requestHttp)

	response, err := requestHttp.Do(ctxt, 5)
	if err != nil {
		return
	}
	logs.Info(response)

	err = json.Unmarshal(response.GetBody(), &res)
	if err != nil {
		return
	}
	return
}
