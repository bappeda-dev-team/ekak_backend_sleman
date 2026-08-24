CREATE TABLE tb_subkegiatan (
    id VARCHAR(225) NOT NULL ,
    rekin_id VARCHAR(225),
    nama_subkegiatan VARCHAR(225),
    tahun VARCHAR(20),
    kode_opd VARCHAR(255),
    pegawai_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE = InnoDB;
