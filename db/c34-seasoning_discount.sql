CREATE TABLE IF NOT EXISTS seasoning_discount (
    discount_id SERIAL REFERENCES discount(discount_id),
    status INT NOT NULL
    discount_percentage INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);