CREATE TABLE
    IF NOT EXISTS label_item (
        product_id SERIAL NOT NULL,
        label_name VARCHAR(255) NOT NULL,
        item_name VARCHAR(255) NOT NULL,
        FOREIGN KEY (product_id, label_name) REFERENCES product_label(product_id, label_name),
        PRIMARY KEY (product_id, label_name, item_name)
    );