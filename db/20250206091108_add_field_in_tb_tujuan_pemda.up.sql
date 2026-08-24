ALTER TABLE tb_tujuan_pemda 
ADD COLUMN tujuan_pemda TEXT,
DROP COLUMN tujuan_pemda_id,
ADD COLUMN tematik_id INT;