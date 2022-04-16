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

	result, err := impl.sintaRP.GetPartnerSinta(ctx, "")
	if err != nil {
		logs.Error("failed get sinta data with error : ", err)
		return res, errors.New(constant.ErrorInternaly)
	}

	res = sintaucm.SintaResponse{
		SintaID:         result.Data.SintaID,
		ScopusID:        result.Data.ScopusID,
		Fullname:        "",
		Nidn:            "",
		GoogleID:        "",
		SintaScore:      "",
		ScopusHindex:    "",
		ScopusCitation:  "",
		ScopusArticle:   "",
		GoogleHindex:    "",
		GoogleCitations: "",
		GoogleArticle:   "",
		GoogleI10:       "",
		RankInNat:       0,
		RankInAffil:     0,
		LastScopusPub:   0,
		Areas:           nil,
	}

	return
}
