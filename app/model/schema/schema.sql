CREATE SCHEMA "pengguna";

DROP TABLE IF EXISTS "identitas_pengguna";
DROP SEQUENCE IF EXISTS identitas_pengguna_seq;
CREATE SEQUENCE identitas_pengguna_seq;
CREATE TABLE "pengguna"."identitas_pengguna"
(
    "id_personal"      bigint    DEFAULT nextval('identitas_pengguna_seq') NOT NULL,
    "id_personal_uuid" bigint                                              NOT NULL,
    "nama_user"        character varying(25)                               NOT NULL,
    "pswd"             character varying(20)                               NOT NULL,
    "tgl_data"         date      default now()                             NOT NULL,
    "kd_sts_pengguna"  character varying(1)                                NOT NULL,
    "id_institusi"     bigint                                              NOT NULL,
    "tgl_updated"      timestamp DEFAULT current_timestamp                 NOT NULL,
    "tgl_created"      timestamp
);
CREATE INDEX identitas_pengguna_id_personal_idx ON pengguna.identitas_pengguna (id_personal);