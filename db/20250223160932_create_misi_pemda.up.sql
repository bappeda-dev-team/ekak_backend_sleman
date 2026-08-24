CREATE TABLE tb_misi_pemda (
    id INT PRIMARY KEY,
    id_visi INT NOT NULL,
    misi TEXT,
    urutan INT NOT NULL,
    tahun_awal_periode VARCHAR(255) NOT NULL,
    tahun_akhir_periode VARCHAR(255) NOT NULL,
    jenis_periode VARCHAR(20) NOT NULL,
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);