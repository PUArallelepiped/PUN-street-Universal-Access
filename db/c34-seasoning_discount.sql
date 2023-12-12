CREATE TABLE
    IF NOT EXISTS seasoning_discount (
        discount_id SERIAL UNIQUE REFERENCES discounts(discount_id),
        discount_percentage INT NOT NULL,
        start_date DATE NOT NULL,
        end_date DATE NOT NULL
    );