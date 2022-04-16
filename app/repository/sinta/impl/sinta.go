package impl

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"project_we/app/model"
	sintarpm "project_we/app/repository/sinta/model"
)

const (
	baseQueryGetSinta = `
							select sinta_id,
								   id_personal,
								   id_sinta,
								   rangking_nasional,
								   rangking_afiliasi,
								   skor_sinta,
								   jml_artikel_google_scholar,
								   jml_sitasi_google_scholar,
								   hindex_google_scholar,
								   i_10_hindex_google_scholar,
								   id_google_scholar,
								   hindex,
								   jml_dokumen,
								   jml_sitasi,
								   tgl_created,
								   tgl_updated
							from sinta
						`
)

func (impl *sintaRepository) GetPartnerSinta(ctx *context.Context, nidn string) (res sintarpm.SintaResponse, err error) {
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

func (impl *sintaRepository) GetSintaByPersonalID(ctx *context.Context, personalID int64) (res []model.Sinta, err error) {
	_, err = impl.orm.QueryTable(new(model.Sinta)).Filter("id_personal", personalID).All(&res)
	return
}

func (impl *sintaRepository) InsertSinta(ctx *context.Context, payload model.Sinta) (err error) {
	dataPayload := new(model.Sinta)
	dataPayload = &payload
	_, err = impl.orm.Insert(dataPayload)
	return
}

func (impl *sintaRepository) UpdateSinta(ctx *context.Context, payload model.Sinta) (err error) {
	data := new(model.Sinta)
	data = &payload
	_, err = impl.orm.Update(data)
	return
}
