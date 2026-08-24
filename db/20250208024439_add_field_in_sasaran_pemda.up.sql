ALTER TABLE tb_sasaran_pemda 
ADD COLUMN sasaran_pemda TEXT,
DROP COLUMN sasaran_pemda_id,
ADD COLUMN subtema_id INT;