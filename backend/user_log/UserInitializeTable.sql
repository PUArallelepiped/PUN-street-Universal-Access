CREATE TABLE IF NOT EXISTS user_datas(
  user_id SERIAL PRIMARY KEY,
  name VARCHAR(32) NOT NULL,
  password VARCHAR(32) NOT NULL,
  email VARCHAR(64) UNIQUE NOT NULL,
  address VARCHAR(64) NOT NULL,
  phone_number VARCHAR(16) NOT NULL,
  birthday DATE,
  authority INT NOT NULL
);