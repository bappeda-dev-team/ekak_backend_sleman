CREATE TABLE tb_usulan_mandatori (
    id VARCHAR(255) NOT NULL,
    usulan TEXT,
    peraturan_terkait TEXT,
    uraian TEXT,
    tahun VARCHAR(20),
    rekin_id VARCHAR(255),
    pegawai_id VARCHAR(255),
    kode_opd VARCHAR(255),
    is_active BOOLEAN DEFAULT FALSE,
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;