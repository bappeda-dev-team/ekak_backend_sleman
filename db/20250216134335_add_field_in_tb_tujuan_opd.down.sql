ALTER TABLE tb_tujuan_opd 
DROP COLUMN periode_id,
CHANGE COLUMN kode_bidang_urusan bidang_urusan_id VARCHAR(255);