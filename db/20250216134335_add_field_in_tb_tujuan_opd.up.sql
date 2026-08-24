ALTER TABLE tb_tujuan_opd 
ADD COLUMN periode_id INT,
CHANGE COLUMN bidang_urusan_id kode_bidang_urusan VARCHAR(255);