package impl

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/pkg/errors"
	sintaucm "project_we/app/usecase/sinta/model"
)

func (impl *sintaUsecase) SyncSinta(ctx *context.Context) (res sintaucm.SintaResponse, errs errors.IError) {
	logs.Info("get sinta data")

	result, err := impl.sintaRP.GetPartnerSinta(ctx, "0315117102")
	if err != nil {
		logs.Error("failed get sinta data with error : ", err)
		return res, errors.New(constant.ErrorInternaly)
	}

	res = sintaucm.SintaResponse{
		SintaID:         result.Data.SintaID,
		ScopusID:        result.Data.ScopusID,
		Fullname:        result.Data.Fullname,
		Nidn:            result.Data.Nidn,
		GoogleID:        result.Data.GoogleID,
		SintaScore:      result.Data.SintaScore,
		ScopusHindex:    result.Data.ScopusHindex,
		ScopusCitation:  result.Data.ScopusCitation,
		ScopusArticle:   result.Data.ScopusArticle,
		GoogleHindex:    result.Data.GoogleHindex,
		GoogleCitations: result.Data.GoogleCitations,
		GoogleArticle:   result.Data.GoogleArticle,
		GoogleI10:       result.Data.GoogleI10,
		RankInNat:       result.Data.RankInNat,
		RankInAffil:     result.Data.RankInAffil,
		LastScopusPub:   result.Data.LastScopusPub,
		Areas:           result.Data.Areas,
	}

	return
}
