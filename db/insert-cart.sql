INSERT INTO cart (product_quantity, customer_id, product_id, store_id) VALUES
    ( 10, (SELECT id from UserData LIMIT 1), (SELECT product_id from product LIMIT 1), (SELECT store_id from stores LIMIT 1) );
