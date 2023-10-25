INSERT INTO products VALUES
    ((default), (SELECT store_id FROM stores WHERE store_id=1), 'pasta', 'tasty pasta', 10, decode('DEADBEEF', 'hex'), 100,1),
    ((default), (SELECT store_id FROM stores WHERE store_id=1), 'black tea', 'student card get free', 20, decode('DEADBEEF', 'hex'), 0,100),
    ((default), (SELECT store_id FROM stores WHERE store_id=2), 'special1', 'special', 10, decode('DEADBEEF', 'hex'), 250,2),
    ((default), (SELECT store_id FROM stores WHERE store_id=2), 'special2', 'special', 10, decode('DEADBEEF', 'hex'), 350,2),
    ((default), (SELECT store_id FROM stores WHERE store_id=2), 'special3', 'special', 10, decode('DEADBEEF', 'hex'), 450,2),
    ((default), (SELECT store_id FROM stores WHERE store_id=3), 'black tea', 'most famous', 20, decode('DEADBEEF', 'hex'), 300,3);
