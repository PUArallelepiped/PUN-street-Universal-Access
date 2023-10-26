CREATE TABLE IF NOT EXISTS shipping_discount (
    discount_id SERIAL REFERENCES discount(discount_id),
    max_price INT NOT NULL,
    store_id SERIAL REFERENCES stores(store_id)
);