package model

import "time"

type (
	UserResponse struct {
		Nama                       string                     `json:"nama,omitempty"`
		NIDN                       string                     `json:"nidn,omitempty"`
		Institusi                  Institusi                  `json:"institusi"`
		ProgramStudi               ProgramStudi               `json:"program_studi"`
		JenjangPendidikanTertinggi JenjangPendidikanTertinggi `json:"jenjang_pendidikan_tertinggi"`
		Alamat                     string                     `json:"alamat,omitempty"`
		TempatLahir                string                     `json:"tempat_lahir,omitempty"`
		TanggalLahir               time.Time                  `json:"tanggal_lahir"`
		NomorKTP                   string                     `json:"nomor_ktp,omitempty"`
		NomorTelepon               string                     `json:"nomor_telepon,omitempty"`
		NomorHP                    string                     `json:"nomor_hp,omitempty"`
		Surel                      string                     `json:"surel,omitempty"`
		WebsitePersonal            string                     `json:"website_personal,omitempty"`
		JabatanFungsional          JabatanFungsional          `json:"jabatan_fungsional"`
	}

	Institusi struct {
		IdInstitusi   int64  `json:"id_institusi,omitempty"`
		NamaInstitusi string `json:"nama_institusi,omitempty"`
	}

	JenjangPendidikanTertinggi struct {
		IdJenjangPendidikanTertinggi   int64  `json:"id_jenjang_pendidikan_tertinggi,omitempty"`
		NamaJenjangPendidikanTertinggi string `json:"nama_jenjang_pendidikan_tertinggi,omitempty"`
	}

	ProgramStudi struct {
		IdProgramStudi   int64  `json:"id_program_studi,omitempty"`
		NamaProgramStudi string `json:"nama_program_studi,omitempty"`
	}

	JabatanFungsional struct {
		IdJabatanFungsional int    `json:"id_jabatan_fungsional,omitempty"`
		JabatanFungsional   string `json:"jabatan_fungsional,omitempty"`
	}
)
