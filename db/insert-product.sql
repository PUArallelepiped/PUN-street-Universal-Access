
INSERT INTO product VALUES
    ((default), (SELECT store_id from stores LIMIT 1), 'name', 'good food', 10, decode('DEADBEEF', 'hex'), 123,2);