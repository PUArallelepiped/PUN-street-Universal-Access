CREATE TABLE
    IF NOT EXISTS user_data(
        user_id SERIAL PRIMARY KEY,
        name VARCHAR(32) NOT NULL,
        password VARCHAR(32) NOT NULL,
        email VARCHAR(64) UNIQUE NOT NULL,
        address VARCHAR(64) NOT NULL,
        phone_number VARCHAR(16) NOT NULL,
        birthday DATE NOT NULL,
        authority VARCHAR(16) NOT NULL,
        current_cart_id INTEGER NOT NULL,
        status INTEGER NOT NULL
    );
