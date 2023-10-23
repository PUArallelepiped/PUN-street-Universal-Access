INSERT INTO carts (product_quantity, customer_id, product_id, store_id) VALUES
    ( 10, (SELECT user_id FROM user_datas WHERE user_id=1 ), (SELECT product_id FROM products WHERE product_id=1), (SELECT store_id FROM stores WHERE store_id=1) ),
    ( 20, (SELECT user_id FROM user_datas WHERE user_id=2 ), (SELECT product_id FROM products WHERE product_id=1), (SELECT store_id FROM stores WHERE store_id=1) ),
    ( 30, (SELECT user_id FROM user_datas WHERE user_id=3 ), (SELECT product_id FROM products WHERE product_id=1), (SELECT store_id FROM stores WHERE store_id=1) );
