CREATE TALBE IF NOT EXISTS discount (
    discount_id SERIAL PRIMARY KEY,
    discount_type INT NOT NULL,
    status INT NOT NULL
    description TEXT NOT NULL,
    name VARCHAR(255) NOT NULL,
);