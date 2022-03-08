-- status_pengguna
INSERT INTO pengguna.status_pengguna (kd_sts_pengguna, status_pengguna, tgl_updated, tgl_created) VALUES (0, 'Off', now(), now());
INSERT INTO pengguna.status_pengguna (kd_sts_pengguna, status_pengguna, tgl_updated, tgl_created) VALUES (1, 'On', now(), now());

-- pengguna
INSERT INTO pengguna.identitas_pengguna (id_personal_uuid, nama_user, pswd, tgl_data, kd_sts_pengguna, id_institusi, tgl_updated, tgl_created) VALUES (1, 'admin', 'admina', now(), '1', 1, now(), now());
INSERT INTO pengguna.identitas_pengguna (id_personal_uuid, nama_user, pswd, tgl_data, kd_sts_pengguna, id_institusi, tgl_updated, tgl_created) VALUES (2, 'dosen', 'dosena', now(), '1', 1, now(), now());

-- peran
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (1, 1, 'Administrator', 'Administrator aplikasi SIM-Litabmas', 'OK', '2016-02-12 11:07:27.433000', '2016-02-12 11:07:27.433000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (2, 1, 'Manajemen Risbang', 'User dng level manajemen di Dit. litabmas', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (4, 1, 'Reviewer Nasional', 'User dengan level reviewer tingkat nasional', 'RK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (5, 1, 'Manajemen PT - Pimpinan', 'User dng level manajemen di perguruan tinggi/kopertis', 'MP', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (6, 1, 'Opt. PT - Penelitian', 'User dengan level operator di perguruan tinggi untuk bidang Penelitian', 'OP', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (7, 1, 'Reviewer PT/Kopertis', 'User dengan level reviewer di tingkat perguruan tinggi/kopertis', 'RP', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (9, 1, 'Pimpinan LP/LPPM - Penelitian', 'User dengan level Pimpinan LP/LPPM - Penelitian', 'MP', '2018-03-29 19:04:17.499000', '2018-03-29 19:04:17.499000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (10, 1, 'Pimpinan LP/LPPM - Pengabdian', 'User dengan level Pimpinan LP/LPPM - Pengabdian', 'MP', '2018-03-29 19:05:13.152000', '2018-03-29 19:05:13.152000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (35, 1, 'Verifikator TKT', 'User dengan level Verifikator TKT', 'VK', '2017-03-26 00:28:26.387000', '2017-03-26 00:28:26.387000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (36, 1, 'Verifikator Akun Non Dosen', 'User dengan level verifikator untuk Non Dosen', 'VI', '2016-06-23 21:46:10.258000', '2016-06-23 21:46:10.258000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (37, 1, 'Pengusul - Dosen', 'User dengan level pengusul kegiatan', 'DP', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (38, 1, 'Operator Risbang', 'User dengan level operator di Dit. Litabmas', 'OK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (39, 1, 'Pengusul - Non Dosen', 'User dengan level Pengusul kegiatan', 'PU', '2016-07-25 14:08:08.853000', '2016-07-25 14:08:08.853000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (40, 1, 'Opt. PT - Pengabdian kpd Masyarakat', 'User dengan level operator di perguruan tinggi untuk bidang Pengabdian kepada Masyarakat', 'OP', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (49, 1, 'Assessor Kinerja Lembaga', 'User dengan level assessor untuk kinerja lembaga', 'VK', '2017-10-06 08:08:29.102000', '2017-10-06 08:08:29.102000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (50, 1, 'Opt. LLDIKTI - Penelitian', 'User dengan level operator di kopertirs untuk bidang Penelitian', 'OC', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (51, 1, 'Opt. LLDIKTI - Pengabdian kpd Masyarakat', 'User dengan level operator di kopertis untuk bidang Pengabdian kepada Masyarakat', 'OC', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (53, 1, 'Super Administrator', 'Super Administrator', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (54, 1, 'Manajemen Risbang - Kasubdit', 'User dng level Manajemen kasubdit di Dit. litabmas', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (55, 1, 'Operator Risbang - Penelitian', 'User dng level operator di Subdit. Penelitian', 'OK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (56, 1, 'Operator Risbang - Abdimas', 'User dng level operator di Subdit. PPM & PKM Seksi Pengabdian', 'OK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (57, 1, 'IT Support', 'User dengan level IT Support', 'OK', '2017-10-06 11:46:09.954000', '2017-10-06 11:46:09.954000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (58, 1, 'Operator Risbang - PE', 'User dng level operator di Subdit. Program dan Evaluasi', 'OK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (59, 1, 'Operator Risbang - HKIP', 'User dng level operator di Subdit. HKIP', 'OK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (60, 1, 'Reviewer BSLN/Insentif Artikel', 'User dengan level reviewer bsln', 'RK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (61, 1, 'Manajemen Risbang - Kasubdit HKIP', 'User dng level Manajemen kasubdit HKI di Dit. litabmas', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (62, 1, 'Manajemen Risbang - Kasubdit Penelitian', 'User dng level Manajemen kasubdit Penelitian di Dit. litabmas', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (63, 1, 'Manajemen Risbang - Kasubdit PPM', 'User dng level Manajemen kasubdit PPM & PKM di Dit. litabmas', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (64, 1, 'Manajemen Risbang - Kasubdit PE', 'User dng level Manajemen Program dan Evaluasi di Dit. litabmas', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (65, 1, 'Manajemen LPNK', 'User dgn level manajemen di LIPI', 'KL', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (66, 1, 'Operator LPNK', 'User dengan level Penanggung Jawab Lembaga - Validator TKT', 'OL', '2017-03-26 00:30:23.688000', '2017-03-26 00:30:23.688000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (67, 1, 'Operator Kementerian Lainnya', 'User dengan level Penanggung Jawab Ristekdikti - Validator TKT', 'OK', '2017-03-26 00:31:34.307000', '2017-03-26 00:31:34.307000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (70, 1, 'Manajemen Kopertis - Pimpinan', 'User dng level manajemen di kopertis', 'MC', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (74, 1, 'Direktur Jenderal Pendidikan Tinggi', 'User dng level Manajemen Dikti', 'MK', '2014-06-06 12:40:45.000000', '2014-06-06 12:40:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (76, 1, 'Operator Pengolah Data', 'User dengan level operator di Dit. Litabmas', 'OK', '2014-07-05 14:37:51.000000', '2014-07-05 14:37:51.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (78, 1, 'Opt. Risbang - Kasi Penelitian', 'User dengan level operator di SubDit. Penelitian', 'OK', '2014-07-12 15:04:16.000000', '2014-07-12 15:05:46.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (79, 1, 'Opt. Risbang - Kasi Pengabdian', 'User dengan level operator di SubDit. KPM', 'MK', '2014-07-12 15:05:25.000000', '2014-07-12 15:05:25.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (81, 1, 'Opt. Risbang - Kasi PE', 'User dengan level operator di SubDit. PE', 'MK', '2014-07-12 15:07:04.000000', '2014-07-12 15:07:04.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (82, 1, 'Opt. Risbang - Kasi HKIP', 'User dengan level operator di SubDit. HKIP', 'OK', '2014-07-13 12:13:50.000000', '2014-07-13 12:13:50.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (86, 1, 'Verifikator Kinerja Penelitian', 'User dengan level verifikator untuk kinerja penelitian', 'VK', '2015-11-25 11:59:28.000000', '2015-11-25 11:59:45.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (87, 1, 'Reviewer Seleksi Administrasi', 'User dengan level verifikator Seleksi Administrasi', 'VK', '2016-10-07 16:16:21.428000', '2016-10-07 16:16:21.428000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (88, 1, 'Verifikator Kinerja Pengabdian', 'User dengan level verifikator untuk kinerja pengabdian', 'VK', '2016-11-14 14:13:05.000000', '2016-11-14 14:13:05.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (89, 1, 'Monitoring Kinerja Pengabdian', 'User dengan level Monitoring untuk Kinerja Pengabdian', 'MK', '2017-06-14 08:20:23.000000', '2017-06-14 08:20:23.000000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (90, 1, 'BPK / Irjen / Bappenas / DJA', 'User dengan level monitoring', 'KL', '2017-03-20 08:57:26.399000', '2017-03-20 08:57:26.399000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (91, 1, 'Reviewer Bantuan Konferensi', 'User dengan level Penilai Bantuan Konferensi Internasional', 'VK', '2017-11-20 16:47:42.819000', '2017-11-20 16:47:42.819000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (92, 1, 'Reviewer Insentif Buku Ajar', 'User dengan level Penilai Insentif Buku Ajar', 'VK', '2018-02-28 05:54:29.131000', '2018-02-28 05:54:29.131000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (93, 1, 'Verifikator Kinerja GESI', 'User dengan  peran Verifikator Kinerja GESI', 'VK', '2018-06-03 15:15:18.413000', '2019-07-25 11:51:34.433000');
INSERT INTO pengguna.peran (id_peran, kd_aplikasi, nama_peran, keterangan, kd_kelompok_peran, tgl_updated, tgl_created) VALUES (94, 1, 'Verifikator Reviewer', 'User dengan peran Verifikator Pendaftaran Reviewer', 'VK', '2019-09-24 12:57:49.245000', '2019-09-24 12:57:49.245000');

-- peran pengguna
INSERT INTO pengguna.peran_pengguna (id_personal, id_peran, kd_sts_peran_pengguna, is_default, is_peran_digunakan, tgl_updated, tgl_created) VALUES (1, 1, '1', 1, 1, now(), now());
INSERT INTO pengguna.peran_pengguna (id_personal, id_peran, kd_sts_peran_pengguna, is_default, is_peran_digunakan, tgl_updated, tgl_created) VALUES (1, 2, '1', 1, 1, now(), now());
