ALTER TABLE tb_sasaran_pemda
DROP COLUMN tujuan_pemda_id,
ADD COLUMN rumus_perhitungan VARCHAR(255) NULL,
ADD COLUMN sumber_data VARCHAR(255) NULL;