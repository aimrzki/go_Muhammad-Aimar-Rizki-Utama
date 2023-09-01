# Buat table product_types
CREATE TABLE product_type (
    id INT(11) PRIMARY KEY not null auto_increment,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# Buat tabel operator
CREATE TABLE operators (
    id INT(11) PRIMARY KEY not null auto_increment,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# Buat tabel products
CREATE TABLE products (
    id INT(11) PRIMARY KEY not null auto_increment,
    product_type_id INT(11),
    operator_id INT(11),
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_type_id) REFERENCES product_type(id),
    FOREIGN KEY (operator_id) REFERENCES operators(id)
);

# Buat table product_description
CREATE TABLE product_description (
    id INT(11) PRIMARY KEY not null auto_increment,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# Buat tabel payment_method
CREATE TABLE payment_methods (
    id INT(11) PRIMARY KEY not null auto_increment,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# Buat table users
CREATE TABLE users (
    id INT(11) PRIMARY KEY not null auto_increment,
    name VARCHAR(255),
    status SMALLINT,
    dob DATE,
    gender CHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# Buat tabel transaction
CREATE TABLE transactions (
    id INT(11) PRIMARY KEY not null auto_increment,
    user_id INT(11),
    payment_method_id INT(11),
    status VARCHAR(10),
    total_qty INT(11),
    total_price NUMERIC(25, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

# Buat table transaction_details
CREATE TABLE transaction_details (
    id INT(11) not null auto_increment PRIMARY KEY,
    transaction_id INT(11),
    product_id INT(11),
    status VARCHAR(10),
    qty INT(11),
    price NUMERIC(25, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

# 1A
INSERT INTO operators (name)
VALUES
    ('P0001'),
    ('P0002'),
    ('P0003'),
    ('P0004'),
    ('P0005');

select * from operators;

# 1B
INSERT INTO product_type (name)
VALUES
    ('Model S'),
    ('Model 3'),
    ('Model Y');

select * from product_type;

# 1C
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES
    (1, 3, 'P001', 'Tesla Model S', 1),
    (1, 3, 'P002', 'Tesla Model S P90D ', 1);

select * from products;

# 1D
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES
    (2, 1, 'P003', 'Tesla Model 3 P50D', 1),
    (2, 1, 'P004', 'Tesla Model 3 P60D', 1),
    (2, 1, 'P005', 'Tesla Model 3 P90D', 1);

select * from products;

# 1E
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES
    (3, 4, 'P006', 'Tesla Model Y P50D', 1),
    (3, 4, 'P007', 'Tesla Model Y P60D', 1),
    (3, 4, 'P008', 'Tesla Model Y P90D', 1);

select * from products;

# 1F
INSERT INTO product_description (description)
VALUES
    ('Mobil Tesla Model S Standard with range 450 KM'),
    ('Mobil Tesla Model S Standard with range 900 KM'),
    ('Mobil Tesla Model 3 Standard with range 500 KM'),
    ('Mobil Tesla Model 3 Standard with range 600 KM'),
    ('Mobil Tesla Model 3 Standard with range 900 KM'),
    ('Mobil Tesla Model Y Standard with range 500 KM'),
    ('Mobil Tesla Model Y Standard with range 600 KM'),
    ('Mobil Tesla Model Y Standard with range 900 KM');

select * from product_description;


# 1G
INSERT INTO payment_methods (name, status)
VALUES
    ('Kartu Debit/Kredit', 1),
    ('Virtual Account', 1),
    ('E-Wallet', 1);

select * from products;

# 1H
INSERT INTO users (status, name,dob, gender)
VALUES
    (1,'Aimar Rizki','2001-05-10', 'M'),
    (1,'Gilang Wahyu','2002-11-21', 'M'),
    (1,'Rafi Faiz','2003-02-15', 'M'),
    (1,'Ami','1998-08-30', 'F'),
    (1,'Yangti','2005-04-05', 'F');

select * from users;

# 1i
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES
    (1, 1, 'Complete', 2, 1000000),
    (1, 2, 'Pending', 1, 2500000),
    (1, 3, 'Complete', 3, 3450000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES
    (2, 1, 'Complete', 1, 450000),
    (2, 2, 'Complete', 2, 750000),
    (2, 3, 'Pending', 1, 325000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES
    (3, 1, 'Pending', 2, 2350000),
    (3, 2, 'Complete', 1, 5500000),
    (3, 3, 'Complete', 3, 425000);

select * from transactions;
select * from transaction_details;
-- Modify the transaction_details table
-- Modify the transaction_details table

# 1J

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (1, 1, 'Complete', 1, 1000000),
    (1, 3, 'Complete', 1, 2500000),
    (1,5, 'Complete', 1, 3450000);

select * from transaction_details;

select * from transactions;

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (2, 2, 'Pending', 1, 450000),
    (2, 4, 'Pending', 1, 75000),
    (2, 6, 'Pending', 1, 325000);

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (3, 1, 'Complete', 1, 2350000),
    (3, 4, 'Complete', 2, 5500000),
    (3, 7, 'Complete', 1, 425000);

select * from transaction_details;

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

select * from products;

# 3B
UPDATE transaction_details
SET qty = 3
WHERE product_id = 1;

select * from transaction_details;

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
SELECT t.*, p.name AS product_name, u.name AS user_name,product_id
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
