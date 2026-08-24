CREATE TABLE tb_pohon_kinerja (
    id INT AUTO_INCREMENT PRIMARY KEY,
    parent INT,
    nama_pohon VARCHAR(255),
    jenis_pohon VARCHAR(255),
    level_pohon INT,
    kode_opd VARCHAR(255),
    keterangan TEXT,
    tahun INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;
