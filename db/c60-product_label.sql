CREATE TABLE
    IF NOT EXISTS product_label (
        product_id SERIAL REFERENCES products(product_id),
        label_name VARCHAR(255) UNIQUE NOT NULL,
        required BOOLEAN NOT NULL,
        PRIMARY KEY (product_id, label_name)
    );