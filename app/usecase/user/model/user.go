package model

import "time"

type (
	UserResponse struct {
		Nama                         string
		NIDN                         string
		IdInstitusi                  int64
		IdProgramStudi               int64
		IdJenjangPendidikanTertinggi int64
		Alamat                       string
		TempatLahir                  string
		TanggalLahir                 time.Time
		NomorKTP                     string
		NomorTelepon                 string
		NomorHP                      string
		Surel                        string
		WebsitePersonal              string
	}
)
