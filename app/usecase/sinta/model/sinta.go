package model

type Scopus struct {
	IDScopus   string `json:"id_scopus"`
	Hindex     int    `json:"hindex"`
	JmlDokumen int    `json:"jml_dokumen"`
	JmlSitasi  int    `json:"jml_sitasi"`
}
type GoogleScholar struct {
	IDGoogleScholar         string `json:"id_google_scholar"`
	HindexGoogleScholar     int    `json:"hindex_google_scholar"`
	JmlArtikelGoogleScholar int    `json:"jml_artikel_google_scholar"`
	JmlSitasiGoogleScholar  int    `json:"jml_sitasi_google_scholar"`
	I10HindexGoogleScholar  int    `json:"i_10_hindex_google_scholar"`
}
type SintaResponse struct {
	IDSinta       int64         `json:"id_sinta"`
	SkorSinta     int           `json:"skor_sinta"`
	Scopus        Scopus        `json:"scopus"`
	GoogleScholar GoogleScholar `json:"google_scholar"`
}
