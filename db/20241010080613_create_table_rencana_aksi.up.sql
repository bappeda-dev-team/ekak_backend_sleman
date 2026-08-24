CREATE TABLE tb_rencana_aksi (
    id VARCHAR(255) PRIMARY KEY,
    rencana_kinerja_id VARCHAR(255),
    urutan INT,
    nama_rencana_aksi VARCHAR(255),
    pegawai_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=INNODB;