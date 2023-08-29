# Buat table product_types
CREATE TABLE product_type (
    id INT(11) PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

# Buat tabel operator
CREATE TABLE operators (
    id INT(11) PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

# Buat tabel products
CREATE TABLE products (
    id INT(11) PRIMARY KEY,
    product_type_id INT(11),
    operator_id INT(11),
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (product_type_id) REFERENCES product_type(id),
    FOREIGN KEY (operator_id) REFERENCES operators(id)
);

# Buat table product_description
CREATE TABLE product_description (
    id INT(11) PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

# Buat tabel payment_method
CREATE TABLE payment_methods (
    id INT(11) PRIMARY KEY,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

# Buat table users
CREATE TABLE users (
    id INT(11) PRIMARY KEY,
    status SMALLINT,
    dob DATE,
    gender CHAR(1),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

# Buat tabel transaction
CREATE TABLE transactions (
    id INT(11) PRIMARY KEY,
    user_id INT(11),
    payment_method_id INT(11),
    status VARCHAR(10),
    total_qty INT(11),
    total_price NUMERIC(25, 2),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

# Buat table transaction_details
CREATE TABLE transaction_details (
    id INT(11) PRIMARY KEY,
    transaction_id INT(11),
    product_id INT(11),
    status VARCHAR(10),
    qty INT(11),
    price NUMERIC(25, 2),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

# 1A
INSERT INTO operators (id, name, created_at, updated_at)
VALUES
    (1, 'Operator A', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Operator B', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 'Operator C', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 'Operator D', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 'Operator E', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

# 1B
INSERT INTO product_type (id, name, created_at, updated_at)
VALUES
    (1, 'Model S', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Model 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 'Model Y', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

# 1C
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
    (1, 1, 3, 'P001', 'Tesla Model S', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 1, 3, 'P002', 'Tesla Model 3', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

# 1D
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
    (3, 2, 1, 'P003', 'Tesla Model X', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 2, 1, 'P004', 'Tesla Model Y', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 2, 1, 'P005', 'Tesla Cybertruck', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

# 1E
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
    (6, 3, 4, 'P006', 'Tesla Roadster', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (7, 3, 4, 'P007', 'Tesla Solar Panel', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (8, 3, 4, 'P008', 'Tesla Charging', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

select * from products;

# 1F
INSERT INTO product_description (id, description, created_at, updated_at)
VALUES
    (1, 'Mobil Tesla Model S', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Mobil Tesla Model 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 'Mobil Tesla Model X', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 'Mobil Tesla Model Y', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 'Truk Tesla', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (6, 'Mobil Tesla Tipe Roadster', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (7, 'Solar Panel', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (8, 'Charger', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


# 1G
INSERT INTO payment_methods (id, name, status, created_at, updated_at)
VALUES
    (1, 'Kartu Debit/Kredit', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Virtual Account', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 'E-Wallet', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

# 1H
INSERT INTO users (id, status, dob, gender, created_at, updated_at)
VALUES
    (1, 1, '2001-05-10', 'M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 1, '2002-11-21', 'F', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 1, '2003-02-15', 'M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 1, '1998-08-30', 'F', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 1, '2005-04-05', 'M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

# 1i
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at)
VALUES
    (1, 1, 1, 'Complete', 2, 1000000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 1, 2, 'Pending', 1, 2500000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 1, 3, 'Complete', 3, 3450000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at)
VALUES
    (4, 2, 1, 'Complete', 1, 450000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 2, 2, 'Complete', 2, 750000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (6, 2, 3, 'Pending', 1, 325000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at)
VALUES
    (7, 3, 1, 'Pending', 2, 2350000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (8, 3, 2, 'Complete', 1, 5500000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (9, 3, 3, 'Complete', 3, 425000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Modify the transaction_details table
-- Modify the transaction_details table

# 1J
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, created_at, updated_at)
VALUES
    (1, 1, 'Complete', 1, 1000000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (1, 3, 'Complete', 1, 2500000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (1, 5, 'Complete', 1, 3450000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, created_at, updated_at)
VALUES
    (2, 2, 'Pending', 1, 450000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 4, 'Pending', 1, 75000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 6, 'Pending', 1, 325000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, created_at, updated_at)
VALUES
    (3, 1, 'Complete', 1, 2350000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 4, 'Complete', 2, 5500000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 7, 'Complete', 1, 425000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

# Menambahkan kolom nama ke tabel user
select * from users;

ALTER TABLE users
ADD COLUMN name VARCHAR(255);

UPDATE users
SET name = 'Aimar Rizki'
WHERE id = 1;

# 2A
SELECT id, name
FROM users
WHERE gender = 'M';

# 2B
SELECT *
FROM products
WHERE id = 3;

# 2C
SELECT *
FROM users
WHERE created_at >= CURDATE() - INTERVAL 7 DAY
    AND created_at <= CURDATE() + INTERVAL 1 DAY
    AND name LIKE '%a%';

# 2D
SELECT COUNT(*) AS jumlah_pelanggan_perempuan
FROM users
WHERE gender = 'F';

# 2E
SELECT *
FROM users
ORDER BY name ASC;

# 2f
SELECT *
FROM products
LIMIT 5;

# 3A
UPDATE products
SET name = 'product dummy'
WHERE id = 1;

# 3B
UPDATE transaction_details
SET qty = 3
WHERE product_id = 1;

# 4A
DELETE FROM products
WHERE id = 1;

# 4A
DELETE FROM products
WHERE product_type_id = 1;

# Join, Union, Sub, Query, Function

# Jawaban NO.1
SELECT *
FROM transactions
WHERE user_id = 1

UNION ALL

SELECT *
FROM transactions
WHERE user_id = 2;

# Jawaban NO.2
SELECT user_id, SUM(total_price) AS total_harga
FROM transactions
WHERE user_id = 1;

# Jawaban NO.3
SELECT SUM(td.qty) AS total_qty, SUM(td.price) AS total_price
FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
WHERE p.product_type_id = 2;

# Jawaban No.4
SELECT p.*, pt.name AS product_type_name
FROM products p
JOIN product_type pt ON p.product_type_id = pt.id;

# Jawaban no 5
SELECT t.*, p.name AS product_name, u.name AS user_name
FROM transactions t
JOIN users u ON t.user_id = u.id
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id;

# Jawaban no 6
DELIMITER //

CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transaction_details
    WHERE transaction_id = OLD.id;
END;
//

DELIMITER ;

# Nomor 7
DELIMITER //

CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
    DECLARE trans_id INT(11);
    SET trans_id = OLD.transaction_id;

    UPDATE transactions
    SET total_qty = total_qty - OLD.qty
    WHERE id = trans_id;
END;
//

DELIMITER ;

# Nomor 8
SELECT *
FROM products
WHERE id NOT IN (
    SELECT DISTINCT product_id
    FROM transaction_details
);
