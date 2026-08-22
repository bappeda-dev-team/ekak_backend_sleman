CREATE TABLE tb_tujuan_opd (
    id INT AUTO_INCREMENT PRIMARY KEY,
    kode_opd VARCHAR(255),
    tujuan TEXT NOT NULL,
    urusan_id VARCHAR(255),
    bidang_urusan_id VARCHAR(255),
    rumus_perhitungan TEXT ,
    sumber_data TEXT ,
    tahun_awal VARCHAR(255),
    tahun_akhir VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);