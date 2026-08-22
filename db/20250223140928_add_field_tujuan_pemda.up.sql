ALTER TABLE tb_tujuan_pemda 
ADD COLUMN tahun_awal_periode VARCHAR(255) NOT NULL,
ADD COLUMN tahun_akhir_periode VARCHAR(255) NOT NULL,
ADD COLUMN jenis_periode VARCHAR(255) NOT NULL,
ADD COLUMN id_visi int NOT NULL,
ADD COLUMN id_misi int NOT NULL;