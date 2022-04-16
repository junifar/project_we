CREATE SCHEMA IF NOT EXISTS "pengguna";
CREATE SCHEMA IF NOT EXISTS "tkt";
CREATE SCHEMA IF NOT EXISTS "pdpt";

DROP TABLE IF EXISTS pengguna."identitas_pengguna";
DROP SEQUENCE IF EXISTS pengguna.identitas_pengguna_seq;
CREATE SEQUENCE pengguna.identitas_pengguna_seq;
CREATE TABLE "pengguna"."identitas_pengguna"
(
    "id_personal"      bigint    DEFAULT nextval('pengguna.identitas_pengguna_seq') NOT NULL,
    "id_personal_uuid" bigint,
    "nama_user"        character varying(25)                                        NOT NULL,
    "pswd"             character varying(20)                                        NOT NULL,
    "tgl_data"         date      default now()                                      NOT NULL,
    "kd_sts_pengguna"  character varying(1)                                         NOT NULL,
    "id_institusi"     bigint                                                       NOT NULL,
    "tgl_updated"      timestamp DEFAULT current_timestamp                          NOT NULL,
    "tgl_created"      timestamp DEFAULT current_timestamp                          NOT NULL,
    CONSTRAINT "identitas_pengguna_pkey" PRIMARY KEY (id_personal)
);
CREATE INDEX identitas_pengguna_id_personal_idx ON pengguna.identitas_pengguna (id_personal);

DROP TABLE IF EXISTS tkt."institusi";
DROP SEQUENCE IF EXISTS tkt.institusi_seq;
CREATE SEQUENCE tkt.institusi_seq;
CREATE TABLE tkt."institusi"
(
    "id_institusi"       bigint    DEFAULT nextval('tkt.institusi_seq') NOT NULL,
    "id_institusi_uuid"  bigint,
    "kd_jenis_institusi" character varying(2)                           NOT NULL,
    "nama_institusi"     character varying(150)                         NOT NULL,
    "alamat"             character varying(200),
    "kd_kota"            bigint,
    "kd_pos"             character varying(5),
    "telepon"            character varying(100),
    "fax"                character varying(100),
    "surel"              character varying(100),
    "website"            character varying(100),
    "id_pdpt"            bigint                                         NOT NULL,
    "kd_sts_aktif"       int                                            NOT NULL,
    "level"              int                                            NOT NULL,
    "id_institusi_induk" bigint,
    "token_verifikasi"   varchar(32),
    "tgl_updated"        timestamp DEFAULT current_timestamp            NOT NULL,
    "tgl_created"        timestamp DEFAULT current_timestamp            NOT NULL,
    CONSTRAINT "institusi_pkey" PRIMARY KEY (id_institusi)
);
CREATE INDEX institusi_id_institusi_idx ON tkt.institusi (id_institusi);


DROP TABLE IF EXISTS pengguna."status_pengguna";
CREATE TABLE pengguna."status_pengguna"
(
    "kd_sts_pengguna" bigint                              NOT NULL,
    "status_pengguna" character varying(10)               NOT NULL,
    "tgl_updated"     timestamp DEFAULT current_timestamp NOT NULL,
    "tgl_created"     timestamp DEFAULT current_timestamp NOT NULL,
    CONSTRAINT "status_pengguna_pkey" PRIMARY KEY (kd_sts_pengguna)
);
CREATE INDEX pengguna_kd_status_pengguna_idx ON pengguna.status_pengguna (kd_sts_pengguna);

DROP TABLE IF EXISTS pengguna."peran";
DROP SEQUENCE IF EXISTS pengguna.peran_seq;
CREATE SEQUENCE pengguna.peran_seq;
CREATE TABLE pengguna."peran"
(
    "id_peran"          bigint    DEFAULT nextval('pengguna.peran_seq') NOT NULL,
    "kd_aplikasi"       int                                             NOT NULL,
    "nama_peran"        character varying(50)                           NOT NULL,
    "keterangan"        character varying(255),
    "kd_kelompok_peran" character varying(2)                            NOT NULL,
    "tgl_updated"       timestamp DEFAULT current_timestamp             NOT NULL,
    "tgl_created"       timestamp DEFAULT current_timestamp             NOT NULL,
    CONSTRAINT "peran_pkey" PRIMARY KEY (id_peran)
);
CREATE INDEX peran_id_peran_idx ON pengguna.peran (id_peran);

DROP TABLE IF EXISTS pengguna."peran_pengguna";
DROP SEQUENCE IF EXISTS pengguna.peran_pengguna_seq;
CREATE SEQUENCE pengguna.peran_pengguna_seq;
CREATE TABLE pengguna."peran_pengguna"
(
    "id_personal"           bigint                              NOT NULL,
    "id_peran"              bigint                              NOT NULL,
    "kd_sts_peran_pengguna" character varying(1)                not null,
    "is_default"            int       DEFAULT 0,
    "is_peran_digunakan"    int       DEFAULT 0,
    "tgl_updated"           timestamp DEFAULT current_timestamp NOT NULL,
    "tgl_created"           timestamp DEFAULT current_timestamp NOT NULL,
    CONSTRAINT "peran_pengguna_pkey" PRIMARY KEY (id_personal, id_peran)
);
CREATE INDEX peran_pengguna_id_peran_id_personal_idx ON pengguna.peran_pengguna (id_personal, id_peran);

DROP TABLE IF EXISTS tkt."jenis_institusi";
DROP SEQUENCE IF EXISTS tkt.jenis_institusi_seq;
CREATE SEQUENCE tkt.jenis_institusi_seq;
CREATE TABLE tkt."jenis_institusi"
(
    "kd_jenis_institusi" bigint    DEFAULT nextval('tkt.jenis_institusi_seq') NOT NULL,
    "jenis_institusi"    character varying(100)                               NOT NULL,
    "tgl_updated"        timestamp DEFAULT current_timestamp                  NOT NULL,
    "tgl_created"        timestamp DEFAULT current_timestamp                  NOT NULL,
    CONSTRAINT "jenis_institusi_pkey" PRIMARY KEY (kd_jenis_institusi)
);
CREATE INDEX jenis_insatitusi_kd_jenis_institusi_idx ON tkt.jenis_institusi (kd_jenis_institusi);

DROP TABLE IF EXISTS tkt."personal";
DROP SEQUENCE IF EXISTS tkt.personal_seq;
CREATE SEQUENCE tkt.personal_seq;
CREATE TABLE "tkt"."personal"
(
    "id_personal"      bigint    DEFAULT nextval('tkt.personal_seq') NOT NULL,
    "id_personal_uuid" bigint,
    "nama"             character varying(100),
    "nomor_ktp"        character varying(25),
    "alamat"           character varying(200),
    "tempat_lahir"     character varying(100),
    "tanggal_lahir"    date,
    "nomor_telepon"    character varying(20),
    "nomor_hp"         character varying(20),
    "surel"            character varying(200),
    "website_personal" character varying(100),
    "id_institusi"     bigint,
    "tgl_updated"      timestamp DEFAULT current_timestamp           NOT NULL,
    "tgl_created"      timestamp DEFAULT current_timestamp           NOT NULL,
    CONSTRAINT "personal_pkey" PRIMARY KEY (id_personal)
);
CREATE INDEX personal_id_personal_idx ON tkt.personal (id_personal);

DROP TABLE IF EXISTS pdpt."dosen";
DROP SEQUENCE IF EXISTS pdpt.dosen_seq;
CREATE SEQUENCE pdpt.dosen_seq;
CREATE TABLE pdpt."dosen"
(
    "nidn"                                character(10)                       NOT NULL,
    "golongan"                            character varying(5)                NOT NULL,
    "pangkat"                             character varying(15)               NOT NULL,
    "kd_perguruan_tinggi"                 character(5)                        NOT NULL,
    "id_program_studi"                    bigint                              NOT NULL,
    "id_fakultas"                         bigint                              NOT NULL,
    "id_jurusan"                          bigint                              NOT NULL,
    "kd_sts_aktif"                        character(1)                        NOT NULL,
    "kd_jenjang_pendidikan_program_studi" character(1)                        NOT NULL,
    "id_personal"                         bigint                              NOT NULL,
    "id_jenjang_pendidikan_tertinggi"     bigint                              NOT NULL,
    "no_sertifikat_dosen"                 character varying(15)               NOT NULL,
    "id_jabatan_fungsional"               int                                 NOT NULL,
    "no_pegawai"                          character varying(20)               NOT NULL,
    "id_pdpt"                             bigint                              NOT NULL,
    "tgl_updated"                         timestamp DEFAULT current_timestamp NOT NULL,
    "tgl_created"                         timestamp DEFAULT current_timestamp NOT NULL,
    CONSTRAINT "dosen_pkey" PRIMARY KEY (nidn)
);
CREATE INDEX nidn_dosen_idx ON pdpt.dosen (nidn);

DROP TABLE IF EXISTS tkt.institusi;
DROP SEQUENCE IF EXISTS tkt.institusi_seq;
CREATE SEQUENCE tkt.institusi_seq;
CREATE TABLE tkt.institusi
(
    id_institusi       bigint DEFAULT nextval('tkt.institusi_seq') NOT NULL,
    kd_jenis_institusi char(2),
    nama_institusi     varchar(150),
    alamat             varchar(200),
    kd_kota            char(5),
    kd_pos             char(5),
    telepon            varchar(100),
    fax                varchar(100),
    surel              varchar(100),
    website            varchar(100),
    id_pdpt            bigint,
    tgl_created        timestamp,
    tgl_updated        timestamp,
    kd_sts_aktif       char,
    level              smallint,
    id_institusi_induk bigint,
    token_verifikasi   varchar(32),
    CONSTRAINT "institusi_pkey" PRIMARY KEY (id_institusi)
);
CREATE INDEX id_institusi_institusi_idx ON tkt.institusi (id_institusi);

DROP TABLE IF EXISTS pdpt.program_studi;
DROP SEQUENCE IF EXISTS pdpt.program_studi_seq;
CREATE SEQUENCE pdpt.program_studi_seq;
CREATE TABLE pdpt.program_studi
(
    id_program_studi      bigint    DEFAULT nextval('pdpt.program_studi_seq') NOT NULL,
    kd_program_studi      char(6)                                             not null,
    kd_perguruan_tinggi   char(6)                                             not null,
    nama_program_studi    varchar(100),
    tgl_created           timestamp default now()                             not null,
    tgl_updated           timestamp default now()                             not null,
    kd_sts_aktif          char      default 1                                 not null,
    kd_jenjang_pendidikan integer,
    id_program_studi_pdpt bigint,
    kd_program_pendidikan integer,
    CONSTRAINT "program_study_pkey" PRIMARY KEY (id_program_studi)
);
create index program_studi_id_program_studi_idx on pdpt.program_studi (id_program_studi);
create index program_studi_kode_perguruan_tinggi_idx on pdpt.program_studi (kd_perguruan_tinggi);

DROP TABLE IF EXISTS pdpt.jenjang_pendidikan;
DROP SEQUENCE IF EXISTS pdpt.jenjang_pendidikan_seq;
CREATE SEQUENCE pdpt.jenjang_pendidikan_seq;
create table pdpt.jenjang_pendidikan
(
    id_jenjang_pendidikan bigint    DEFAULT nextval('pdpt.jenjang_pendidikan_seq') NOT NULL,
    jenjang_pendidikan    varchar(50),
    degree                varchar(50),
    tgl_created           timestamp default now()                                  not null,
    tgl_updated           timestamp default now()                                  not null,
    kd_sts_aktif          char      default 1                                      not null,
    id_tingkat_pendidikan integer,
    CONSTRAINT "jenjang_pendidikan_pkey" PRIMARY KEY (id_jenjang_pendidikan)
);
create index if not exists jenjang_pendidikan_kode_jenjang_pendidikan_idx on pdpt.jenjang_pendidikan (id_jenjang_pendidikan);

DROP TABLE IF EXISTS pdpt.jabatan_fungsional;
DROP SEQUENCE IF EXISTS pdpt.jabatan_fungsional_seq;
CREATE SEQUENCE pdpt.jabatan_fungsional_seq;
create table pdpt.jabatan_fungsional
(
    id_jabatan_fungsional bigint    DEFAULT nextval('pdpt.jabatan_fungsional_seq') NOT NULL,
    jabatan_fungsional    varchar(50),
    tgl_created           timestamp default now()                                  not null,
    tgl_updated           timestamp default now()                                  not null,
    kd_sts_aktif          char      default 1                                      not null,
    CONSTRAINT "jabatan_fungsional_pkey" PRIMARY KEY (id_jabatan_fungsional)
);
create index if not exists jabatan_fungsional_id_jabatan_fungsional_idx on pdpt.jabatan_fungsional (id_jabatan_fungsional);

DROP TABLE IF EXISTS sinta;
DROP SEQUENCE IF EXISTS sinta_seq;
CREATE SEQUENCE sinta_seq;
create table if not exists sinta
(
    sinta_id                   bigint    DEFAULT nextval('sinta_seq') NOT NULL,
    id_personal                bigint                                 not null,
    id_sinta                   bigint                                 not null,
    skor_sinta                 integer,
    jml_artikel_google_scholar integer,
    jml_sitasi_google_scholar  integer,
    hindex_google_scholar      integer,
    i_10_hindex_google_scholar integer,
    id_google_scholar          varchar(50),
    hindex                     integer,
    jml_dokumen                integer,
    jml_sitasi                 integer,
    tgl_created                timestamp default now()                not null,
    tgl_updated                timestamp default now()                not null,
    CONSTRAINT "sinta_pkey" PRIMARY KEY (sinta_id)
);
create index if not exists sinta_id_personal_idx on sinta (id_personal);