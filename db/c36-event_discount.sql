CREATE TABLE IF NOT EXISTS event_discount (
    discount_id SERIAL REFERENCES discount(discount_id),
    max_quantity INT NOT NULL,
    product_id SERIAL REFERENCES products(product_id),
);