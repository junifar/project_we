package model

type SintaResponse struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

type Data struct {
	SintaID         string   `json:"sinta_id"`
	ScopusID        string   `json:"scopus_id"`
	Fullname        string   `json:"fullname"`
	Nidn            string   `json:"NIDN"`
	GoogleID        string   `json:"google_id"`
	SintaScore      string   `json:"sinta_score"`
	ScopusHindex    string   `json:"scopus_hindex"`
	ScopusCitation  string   `json:"scopus_citation"`
	ScopusArticle   string   `json:"scopus_article"`
	GoogleHindex    string   `json:"google_hindex"`
	GoogleCitations string   `json:"google_citations"`
	GoogleArticle   string   `json:"google_article"`
	GoogleI10       string   `json:"google_i10"`
	RankInNat       int      `json:"rank_in_nat"`
	RankInAffil     int      `json:"rank_in_affil"`
	LastScopusPub   int      `json:"last_scopus_pub"`
	Areas           []string `json:"areas"`
}
