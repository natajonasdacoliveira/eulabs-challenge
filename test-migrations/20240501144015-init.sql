-- +migrate Up
CREATE TABLE product (
  id BigInt(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  price DECIMAL(10, 2),
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL DEFAULT NULL
);

INSERT INTO
  product (id, name, price, description)
VALUES
  (
    null,
    'Product 1',
    19.99,
    'Sample description for Product 1'
  ),
  (
    null,
    'Product 2',
    29.99,
    'Sample description for Product 2'
  ),
  (
    null,
    'Product 3',
    39.99,
    'Sample description for Product 3'
  );

-- +migrate Down
DROP TABLE product;