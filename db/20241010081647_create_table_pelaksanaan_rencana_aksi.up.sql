CREATE TABLE tb_pelaksanaan_rencana_aksi (
    id VARCHAR(255),
    rencana_aksi_id VARCHAR(255),
    bobot INT,
    bulan INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=INNODB;