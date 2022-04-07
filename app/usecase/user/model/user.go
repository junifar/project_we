package model

import "time"

type (
	UserResponse struct {
		Nama                       string
		NIDN                       string
		Institusi                  Institusi
		IdProgramStudi             int64
		JenjangPendidikanTertinggi JenjangPendidikanTertinggi
		Alamat                     string
		TempatLahir                string
		TanggalLahir               time.Time
		NomorKTP                   string
		NomorTelepon               string
		NomorHP                    string
		Surel                      string
		WebsitePersonal            string
	}

	Institusi struct {
		IdInstitusi   int64
		NamaInstitusi string
	}

	JenjangPendidikanTertinggi struct {
		IdJenjangPendidikanTertinggi   int64
		NamaJenjangPendidikanTertinggi string
	}
)
