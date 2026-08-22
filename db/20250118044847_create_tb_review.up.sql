CREATE TABLE tb_review (
    id INT,
    id_pohon_kinerja INT,
    review TEXT,
    keterangan TEXT,
    created_by VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)ENGINE=InnoDB;
