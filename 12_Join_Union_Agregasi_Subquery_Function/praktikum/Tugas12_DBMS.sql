create table operators(
    id int(11) not null auto_increment,
    name varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key (id)
);

create table product_types(
    id int(11) not null auto_increment,
    name varchar(255),
     created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key (id)
);

create table products(
    id int(11) not null auto_increment,
    product_type_id int(11),
    operator_id int(11),
    code varchar(50),
    name varchar(100),
    status smallint,
    created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key (id),
    foreign key (product_type_id) references product_types(id),
    foreign key (operator_id) references  operators(id)
);

create table product_description (
    id int (11) not null auto_increment ,
    description text,
     created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key(id)
);

create table payment_methods(
    id int(11) not null auto_increment,
    name varchar(255),
    status smallint,
     created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key (id)
);

create table users(
    id int(11) not null auto_increment,
    status smallint,
    dob date,
    gender char(1),
    created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key (id)
);

create table transactions (
    id int(11) not null auto_increment,
    user_id int(11),
    payment_method_id int(11),
    status varchar(10),
    total_qty int(11),
    total_price numeric(25,2),
    created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (payment_method_id) references payment_methods(id)
);

create table transactions_detail(
    transaction_id int(11) not null,
    product_id int(11),
    status varchar(10),
    qty int(11),
    price numeric(25,2),
    created_at timestamp default current_timestamp,
    updated_at timestamp default  current_timestamp,
    primary key (transaction_id,product_id),
    foreign key (transaction_id) references transactions(id),
    foreign key (product_id) references products(id)
);

# 1A

insert into operators (name)
values ('operator A'),
       ('operator B'),
       ('operator C'),
       ('operator D'),
       ('operator E');

# 1B

insert into product_types (name)
values ('Sedan'),
       ('SUV'),
       ('Truck');

# 1C
insert into products (product_type_id, operator_id, code, name, status)
values (1,3,'T0001','Tesla Model S P50D',1),
       (1,3,'T0001','Tesla Model S P90D',1);

# 1D
insert into products (product_type_id, operator_id, code, name, status)
values (2,1,'T0002','Tesla Model X P50D',1),
       (2,1,'T0003','Tesla Model X P90D',1);

# 1E
insert into products (product_type_id, operator_id, code, name, status)
values (3,4,'T0004','Tesla Cybertruck White',0),
       (3,4,'T0005','Tesla Cybertruck Black',0);

# 1F
insert into product_description (description)
values ('Tesla model S with range 500km'),
       ('Tesla model S with range 900km'),
       ('Tesla model X with range 500km'),
       ('Tesla model X with range 900km'),
       ('Tesla Cybertruck with white color'),
       ('Tesla Cybertruck with black color');

select * from product_description;

alter table products
add column  descriptions int(11);

alter table products
add constraint fk_product_description
foreign key (descriptions) references product_description(id);

update products
set descriptions = 1
where id =1;

update products
set descriptions = 2
where id =2;

update products
set descriptions = 3
where id =3;

update products
set descriptions = 4
where id =4;

update products
set descriptions = 5
where id =5;

update products
set descriptions = 6
where id =6;

select * from products
join product_description on (products.descriptions = product_description.id);

# 1G

insert into payment_methods (name, status)
values ('Transfer Bank',1),
       ('Virtual Account',1),
       ('Kartu Debit / Kredit',1),
       ('E-wallet',1),
       ('COD',1);

# 1H

alter table users
add column name varchar(255) not null ;

INSERT INTO users (status, name,dob, gender)
VALUES
    (1,'Aimar Rizki','2001-05-10', 'M'),
    (1,'Gilang Wahyu','2002-11-21', 'M'),
    (1,'Rafi Faiz','2003-02-15', 'M'),
    (1,'Ami','1998-08-30', 'F'),
    (1,'Yangti','2005-04-05', 'F');

# 1I

insert into transactions (user_id, payment_method_id, status, total_qty, total_price)
values (1,3,'success',1,100000),
       (1,2,'success',1,150000),
       (1,1,'success',1,200000),
       (2,2,'success',1,300000),
       (2,3,'success',1,450000),
       (2,4,'success',1,500000),
       (3,5,'success',1,200000),
       (3,5,'success',1,350000),
       (3,5,'success',1,400000),
       (4,2,'success',1,200000),
       (4,1,'success',1,150000),
       (4,1,'success',1,200000),
       (5,3,'success',1,110000),
       (5,5,'success',1,150000),
       (5,4,'success',1,200000);

# 1J

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (1,1,'sucess',2,1000000),
       (1,3,'sucess',1,3000000),
       (1,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (2,2,'sucess',2,1000000),
       (2,3,'sucess',1,3000000),
       (2,5,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (3,1,'sucess',2,1000000),
       (3,3,'sucess',1,3000000),
       (3,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (4,1,'sucess',2,1000000),
       (4,3,'sucess',1,3000000),
       (4,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (5,1,'sucess',2,1000000),
       (5,3,'sucess',1,3000000),
       (5,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (6,1,'sucess',2,1000000),
       (6,3,'sucess',1,3000000),
       (6,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (7,1,'sucess',2,1000000),
       (7,3,'sucess',1,3000000),
       (7,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (8,1,'sucess',2,1000000),
       (8,3,'sucess',1,3000000),
       (8,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (9,1,'sucess',2,1000000),
       (9,3,'sucess',1,3000000),
       (9,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (10,1,'sucess',2,1000000),
       (10,3,'sucess',1,3000000),
       (10,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (11,1,'sucess',2,1000000),
       (11,3,'sucess',1,3000000),
       (11,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (12,1,'sucess',2,1000000),
       (12,3,'sucess',1,3000000),
       (12,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (13,1,'sucess',2,1000000),
       (13,3,'sucess',1,3000000),
       (13,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (14,1,'sucess',2,1000000),
       (14,3,'sucess',1,3000000),
       (14,4,'sucess',3,4000000);

insert into transactions_detail (transaction_id, product_id, status, qty, price)
values (15,1,'sucess',2,1000000),
       (15,3,'sucess',1,3000000),
       (15,4,'sucess',3,4000000);

select * from transactions_detail;

# Select

# 2A
select name from users
where gender='M';

# 2B
select * from products
where id=3;

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

# 2F
select * from products
limit 5;

# Update

# 3A
UPDATE products
set name = 'product dummy'
where id = 1;

# 3B
update transactions_detail
set qty=3
where product_id=1;

# 4A
DELETE FROM products
WHERE id = 1;

# 4B
DELETE FROM products
WHERE product_type_id = 1;

# Join Union Subquery Function

# 1
SELECT *
FROM transactions
WHERE user_id = 1

UNION ALL

SELECT *
FROM transactions
WHERE user_id = 2;

# 2
SELECT user_id, SUM(total_price) AS total_harga
FROM transactions
WHERE user_id = 1;

# 3
select sum(transactions_detail.qty),sum(transactions_detail.price)
from transactions
join transactions_detail on(transactions.id=transactions_detail.transaction_id)
join products on (transactions_detail.product_id=products.id)
where products.product_type_id = 2;

#4

SELECT *,product_types.name
FROM products
JOIN product_types on (products.product_type_id=product_types.id);

# Jawaban no 5

SELECT transactions.*, products.name AS product_name, users.name AS user_name, transactions_detail.product_id
FROM transactions
JOIN users ON transactions.user_id = users.id
JOIN transactions_detail ON transactions.id = transactions_detail.transaction_id
JOIN products ON transactions_detail.product_id = products.id;

# Jawaban no 6

DELIMITER //

CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transactions_detail
    WHERE transaction_id = OLD.id;
END;
//

DELIMITER ;

# Nomor 7
DELIMITER //

CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transactions_detail
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
    FROM transactions_detail
);