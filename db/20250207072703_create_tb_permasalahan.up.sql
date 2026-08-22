CREATE TABLE tb_permasalahan (
    id INT PRIMARY KEY,
    rekin_id VARCHAR(255) ,
    permasalahan VARCHAR(255) NOT NULL,
    penyebab_internal TEXT,
    penyebab_eksternal TEXT,
    jenis_permasalahan VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;
