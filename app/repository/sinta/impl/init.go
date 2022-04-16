package impl

import (
	"github.com/beego/beego/v2/adapter/orm"
	"project_we/app/pkg/curl"

	sintarp "project_we/app/repository/sinta"
)

type sintaRepository struct {
	orm  orm.Ormer
	curl curl.IHttpRequestor
}

// NewSintaRepository ...
func NewSintaRepository(orm orm.Ormer, curl curl.IHttpRequestor) sintarp.ISinta {
	return &sintaRepository{
		orm:  orm,
		curl: curl,
	}
}
