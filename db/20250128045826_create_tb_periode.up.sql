CREATE TABLE tb_periode (
    id INT PRIMARY KEY,
    tahun_awal VARCHAR(50) NOT NULL,
    tahun_akhir VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY tahun_awal_unique (tahun_awal),
    UNIQUE KEY tahun_akhir_unique (tahun_akhir)
) ENGINE=InnoDB;