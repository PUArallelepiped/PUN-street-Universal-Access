CREATE TABLE
    IF NOT EXISTS label_item (
        label_name VARCHAR(255) NOT NULL REFERENCES product_label(label_name),
        item_name VARCHAR(255) NOT NULL,
        PRIMARY KEY (label_name, item_name)
    );