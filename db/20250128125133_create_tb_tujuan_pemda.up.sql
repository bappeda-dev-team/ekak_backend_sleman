CREATE TABLE tb_tujuan_pemda (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tujuan_pemda_id INT UNIQUE,
    periode_id INT,
    rumus_perhitungan VARCHAR(255),
    sumber_data VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB;