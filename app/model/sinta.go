package model

import "time"

import (
	"github.com/beego/beego/v2/client/orm"
)

type Sinta struct {
	SintaID                 int64     `orm:"column(sinta_id);pk;auto"`
	IDPersonal              int64     `orm:"column(id_personal)"`
	IDSinta                 int64     `orm:"column(id_sinta)"`
	IDScopus                string    `orm:"column(id_scopus)"`
	SkorSinta               int       `orm:"column(skor_sinta)"`
	JmlArtikelGoogleScholar int       `orm:"column:(jml_artikel_google_scholar)"`
	JmlSitasiGoogleScholar  int       `orm:"column(jml_sitasi_google_scholar)"`
	HindexGoogleScholar     int       `orm:"column(hindex_google_scholar)"`
	I10HindexGoogleScholar  int       `orm:"column(i_10_hindex_google_scholar)"`
	IDGoogleScholar         string    `orm:"column(id_google_scholar)"`
	Hindex                  int       `orm:"column(hindex)"`
	JmlDokumen              int       `orm:"column(jml_dokumen)"`
	JmlSitasi               int       `orm:"column(jml_sitasi)"`
	TglCreated              time.Time `orm:"column(tgl_created)"`
	TglUpdated              time.Time `orm:"column(tgl_updated)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Sinta))
}
