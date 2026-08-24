ALTER TABLE tb_subkegiatan 
DROP COLUMN status,
DROP COLUMN rekin_id,
ADD COLUMN pegawai_id VARCHAR(255) NOT NULL;