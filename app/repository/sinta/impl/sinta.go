package impl

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"project_we/app/repository/sinta/model"
)

func (impl *sintaRepository) GetPartnerSinta(ctx *context.Context, nidn string) (res model.SintaResponse, err error) {
	ctxt := context2.Background()
	var apiKey string
	apiKey = "a39e735fd5049ba1f7ff0b4e05c9f207"
	url := fmt.Sprintf("https://sinta.kemdikbud.go.id/api/select?nidn=%s&api_key=%s", nidn, apiKey)
	requestHttp := impl.curl.NewHttpRequest(http.MethodPost, url)

	result := map[string]interface{}{}
	response, err := requestHttp.Do(ctxt, 5)
	if err != nil {
		result["error"] = err.Error()
	}
	err = json.Unmarshal(response.GetBody(), &res)
	if err != nil {
		return
	}
	return
}
