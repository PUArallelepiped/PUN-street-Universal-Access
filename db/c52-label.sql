CREATE TABLE
    IF NOT EXISTS labels (
        category_id SERIAL REFERENCES categories(category_id),
        store_id SERIAL REFERENCES stores(store_id),
        PRIMARY KEY (category_id, store_id)
    );