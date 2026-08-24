CREATE TABLE tb_dasar_hukum (
    id VARCHAR(255) NOT NULL,
    rekin_id VARCHAR(255) ,
    urutan INT ,
    peraturan_terkait TEXT,
    uraian TEXT,
    pegawai_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;
