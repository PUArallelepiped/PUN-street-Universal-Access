CREATE TABLE
    IF NOT EXISTS categories (
        category_id SERIAL PRIMARY KEY,
        name VARCHAR(255) UNIQUE NOT NULL
    );