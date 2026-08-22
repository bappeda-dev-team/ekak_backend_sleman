CREATE TABLE tb_inovasi (
    id VARCHAR(255) NOT NULL,
    rekin_id VARCHAR(255) NOT NULL,
    judul_inovasi TEXT,
    jenis_inovasi TEXT,
    gambaran_nilai_kebaruan TEXT,
    pegawai_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;
