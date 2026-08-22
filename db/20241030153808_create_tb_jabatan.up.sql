CREATE TABLE tb_jabatan (
    id VARCHAR(36) NOT NULL,
    nama_jabatan VARCHAR(255) NOT NULL,
    kelas_jabatan VARCHAR(255) NOT NULL,
    jenis_jabatan VARCHAR(255) NOT NULL,
    nilai_jabatan INT DEFAULT 0,
    kode_opd VARCHAR(255) NOT NULL,
    index_jabatan INT DEFAULT 0,
    tahun VARCHAR(255) NOT NULL,
    esselon VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;
