package model

import "time"

type (
	UserResponse struct {
		Nama                       string                     `json:"nama"`
		NIDN                       string                     `json:"nidn"`
		Institusi                  Institusi                  `json:"institusi"`
		ProgramStudi               ProgramStudi               `json:"program_studi"`
		JenjangPendidikanTertinggi JenjangPendidikanTertinggi `json:"jenjang_pendidikan_tertinggi"`
		Alamat                     string                     `json:"alamat"`
		TempatLahir                string                     `json:"tempat_lahir"`
		TanggalLahir               time.Time                  `json:"tanggal_lahir"`
		NomorKTP                   string                     `json:"nomor_ktp"`
		NomorTelepon               string                     `json:"nomor_telepon"`
		NomorHP                    string                     `json:"nomor_hp"`
		Surel                      string                     `json:"surel"`
		WebsitePersonal            string                     `json:"website_personal"`
		JabatanFungsional          JabatanFungsional          `json:"jabatan_fungsional"`
	}

	Institusi struct {
		IdInstitusi   int64  `json:"id_institusi"`
		NamaInstitusi string `json:"nama_institusi"`
	}

	JenjangPendidikanTertinggi struct {
		IdJenjangPendidikanTertinggi   int64  `json:"id_jenjang_pendidikan_tertinggi"`
		NamaJenjangPendidikanTertinggi string `json:"nama_jenjang_pendidikan_tertinggi"`
	}

	ProgramStudi struct {
		IdProgramStudi   int64  `json:"id_program_studi"`
		NamaProgramStudi string `json:"nama_program_studi"`
	}

	JabatanFungsional struct {
		IdJabatanFungsional int    `json:"id_jabatan_fungsional"`
		JabatanFungsional   string `json:"jabatan_fungsional"`
	}
)
