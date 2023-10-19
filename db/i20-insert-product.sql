INSERT INTO products VALUES
    ((default), (SELECT store_id FROM stores WHERE store_id=1), 'name_1', 'good food_1', 10, decode('DEADBEEF', 'hex'), 100,1),
    ((default), (SELECT store_id FROM stores WHERE store_id=2), 'name_2', 'good food_2', 20, decode('DEADBEEF', 'hex'), 200,2),
    ((default), (SELECT store_id FROM stores WHERE store_id=3), 'name_3', 'good food_3', 30, decode('DEADBEEF', 'hex'), 300,3);