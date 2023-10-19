INSERT INTO cart (product_quantity, customer_id, product_id, store_id) VALUES
    ( 10, (SELECT id FROM UserData WHERE id=1 ), (SELECT product_id FROM product WHERE product_id=1), (SELECT store_id FROM stores WHERE store_id=1) );

INSERT INTO cart (product_quantity, customer_id, product_id, store_id) VALUES
    ( 10, (SELECT id FROM UserData WHERE id=2 ), (SELECT product_id FROM product WHERE product_id=1), (SELECT store_id FROM stores WHERE store_id=1) );
    
INSERT INTO cart (product_quantity, customer_id, product_id, store_id) VALUES
    ( 10, (SELECT id FROM UserData WHERE id=3 ), (SELECT product_id FROM product WHERE product_id=1), (SELECT store_id FROM stores WHERE store_id=1) );
