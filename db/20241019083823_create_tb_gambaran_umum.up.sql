CREATE TABLE tb_gambaran_umum (
    id VARCHAR(255) NOT NULL,
    rekin_id VARCHAR(255) ,
    urutan INT,
    gambaran_umum TEXT,
    pegawai_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;
