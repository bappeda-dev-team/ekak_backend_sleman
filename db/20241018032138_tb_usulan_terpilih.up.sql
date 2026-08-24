CREATE TABLE tb_usulan_terpilih (
    id VARCHAR(255) NOT NULL,
    keterangan TEXT,
    jenis_usulan ENUM('mandatori', 'musrebang', 'inisiatif', 'pokok_pikiran') NOT NULL,
    usulan_id VARCHAR(255),
    rekin_id VARCHAR(255),
    tahun VARCHAR(20),
    kode_opd VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;