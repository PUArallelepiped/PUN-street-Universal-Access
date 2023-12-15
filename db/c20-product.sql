CREATE TABLE
    IF NOT EXISTS products (
        product_id SERIAL PRIMARY KEY,
        store_id SERIAL REFERENCES stores(store_id),
        name VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        picture VARCHAR(255) NOT NULL,
        price INTEGER NOT NULL,
        stock INTEGER NOT NULL,
        status INTEGER NOT NULL
    );
