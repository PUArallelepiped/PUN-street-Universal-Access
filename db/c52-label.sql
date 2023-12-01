CREATE TABLE
    IF NOT EXISTS labels (
        category_id SERIAL REFERENCES categorys(category_id),
        store_id SERIAL REFERENCES stores(store_id),
        PRIMARY KEY (category_id, store_id)
    );