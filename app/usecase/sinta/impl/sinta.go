package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/model"
	"project_we/app/pkg/errors"
	sintaucm "project_we/app/usecase/sinta/model"
	"strconv"
	"time"
)

func (impl *sintaUsecase) SyncSinta(ctx *context.Context) (res sintaucm.SintaResponse, errs errors.IError) {
	logs.Info("get sinta data")
	var personalID int64
	personalID = 2

	result, err := impl.sintaRP.GetPartnerSinta(ctx, "0315117102")
	if err != nil {
		logs.Error("failed get sinta data with error : ", err)
		return res, errors.New(constant.ErrorInternaly)
	}

	sintaData, err := impl.sintaRP.GetSintaByPersonalID(ctx, personalID)
	if err != nil {
		logs.Error("failed get sinta by personal id ", personalID, " with error : ", err)
		return res, errors.New(constant.ErrorInternaly)
	}

	sintaID, _ := strconv.ParseInt(result.Data.SintaID, 10, 64)
	skorSinta, _ := strconv.Atoi(result.Data.SintaScore)
	hindex, _ := strconv.Atoi(result.Data.ScopusHindex)
	jmlDokumen, _ := strconv.Atoi(result.Data.ScopusArticle)
	jmlSitasi, _ := strconv.Atoi(result.Data.ScopusCitation)
	hindexGoogleScholar, _ := strconv.Atoi(result.Data.GoogleHindex)
	jmlArtikelGoogleScholar, _ := strconv.Atoi(result.Data.GoogleArticle)
	jmlSitasiGoogleScholar, _ := strconv.Atoi(result.Data.GoogleCitations)
	I10Hindex, _ := strconv.Atoi(result.Data.GoogleI10)

	now := time.Now()
	if len(sintaData) == 0 {
		logs.Info("insert sinta data with personal id : ", personalID)
		sintaPayload := model.Sinta{
			IDPersonal:              personalID,
			IDSinta:                 sintaID,
			SkorSinta:               skorSinta,
			JmlArtikelGoogleScholar: jmlArtikelGoogleScholar,
			JmlSitasiGoogleScholar:  jmlSitasiGoogleScholar,
			HindexGoogleScholar:     hindexGoogleScholar,
			I10HindexGoogleScholar:  I10Hindex,
			IDGoogleScholar:         result.Data.GoogleID,
			Hindex:                  hindex,
			JmlDokumen:              jmlDokumen,
			JmlSitasi:               jmlSitasi,
			TglCreated:              now,
			TglUpdated:              now,
		}
		err := impl.sintaRP.InsertSinta(ctx, sintaPayload)
		if err != nil {
			logs.Error("failed insert sinta with error : ", err)
			return res, errors.New(constant.ErrorInternaly)
		}
	} else {
		logs.Info("update sinta data with personal id : ", personalID)
		sintaData[0].SkorSinta = skorSinta
		sintaData[0].JmlArtikelGoogleScholar = jmlArtikelGoogleScholar
		sintaData[0].JmlSitasiGoogleScholar = jmlSitasiGoogleScholar
		sintaData[0].HindexGoogleScholar = hindexGoogleScholar
		sintaData[0].I10HindexGoogleScholar = I10Hindex
		sintaData[0].Hindex = hindex
		sintaData[0].JmlDokumen = jmlDokumen
		sintaData[0].JmlSitasi = jmlSitasi
		sintaData[0].TglUpdated = now
		err := impl.sintaRP.InsertSinta(ctx, sintaData[0])
		if err != nil {
			logs.Error("failed update sinta with error : ", err)
			return res, errors.New(constant.ErrorInternaly)
		}
	}

	res = sintaucm.SintaResponse{
		IDSinta:   sintaID,
		SkorSinta: skorSinta,
		Scopus: sintaucm.Scopus{
			IDScopus:   result.Data.ScopusID,
			Hindex:     hindex,
			JmlDokumen: jmlDokumen,
			JmlSitasi:  jmlSitasi,
		},
		GoogleScholar: sintaucm.GoogleScholar{
			IDGoogleScholar:         result.Data.GoogleID,
			HindexGoogleScholar:     hindexGoogleScholar,
			JmlArtikelGoogleScholar: jmlArtikelGoogleScholar,
			JmlSitasiGoogleScholar:  jmlSitasiGoogleScholar,
			I10HindexGoogleScholar:  I10Hindex,
		},
	}

	return
}
