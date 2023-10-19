INSERT INTO products VALUES
    ((default), (SELECT store_id FROM stores WHERE store_id=1), 'name', 'good food', 10, decode('DEADBEEF', 'hex'), 123,2),
    ((default), (SELECT store_id FROM stores WHERE store_id=2), 'name', 'good food', 10, decode('DEADBEEF', 'hex'), 123,2),
    ((default), (SELECT store_id FROM stores WHERE store_id=3), 'name', 'good food', 10, decode('DEADBEEF', 'hex'), 123,2);