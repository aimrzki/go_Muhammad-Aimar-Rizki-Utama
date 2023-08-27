-- Jawaban Nomor 2a --
CREATE TABLE user (
    id INT PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

-- Jawaban Nomor 2b --
CREATE TABLE product (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    price DECIMAL(10, 2),
    product_type_id INT,
    operator_id INT,
    description_id INT,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE product_type (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE operators (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE product_description (
    id INT PRIMARY KEY,
    description TEXT
);

CREATE TABLE payment_method (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    description_id INT,
    created_at DATETIME,
    updated_at DATETIME
);

-- Jawaban Nomor 2c --
CREATE TABLE payment_method_description (
    id INT PRIMARY KEY,
    description TEXT
);

CREATE TABLE transaction (
    id INT PRIMARY KEY,
    user_id INT,
    total_amount DECIMAL(10, 2),
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE transaction_detail (
    id INT PRIMARY KEY,
    transaction_id INT,
    product_id INT,
    quantity INT,
    subtotal DECIMAL(10, 2),
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (transaction_id) REFERENCES transaction(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);

-- Jawaban Nomor 3 --
CREATE TABLE kurir (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

-- Jawaban Nomor 4 --
ALTER TABLE kurir
ADD COLUMN ongkos_dasar DECIMAL(10, 2);

-- Jawaban Nomor 5 --
ALTER TABLE kurir
RENAME TO shipping;

-- Jawaban Nomor 6 --
DROP TABLE shipping;

-- Jawaban Nomor 7a sudah di jawab di atas --


-- Jawaban Nomor 7b --
CREATE TABLE alamat (
    id INT PRIMARY KEY,
    user_id INT,
    address TEXT,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(id)
);


-- Jawaban Nomor 7c --
CREATE TABLE user_payment_method_detail (
    id INT PRIMARY KEY,
    user_id INT,
    payment_method_id INT,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);

show tables;
