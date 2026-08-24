CREATE TABLE tb_rencana_kinerja (
    id VARCHAR(255) NOT NULL UNIQUE,
    nama_rencana_kinerja VARCHAR(255) NOT NULL,
    tahun VARCHAR(20) ,
    status_rencana_kinerja VARCHAR(255) ,
    catatan text ,
    pegawai_id VARCHAR(255),
    kode_opd VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;