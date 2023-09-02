# Membuat tabel user
create table user (
    id int not null auto_increment,
    nama varchar(255) not null ,
    alamat varchar(255) not null ,
    tanggal_lahir varchar(255) not null ,
    status varchar(255) not null ,
    gender enum ('pria','wanita'),
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    primary key (id)
);

# Insert tabel user
insert into user (nama, alamat, tanggal_lahir, status, gender)
VALUES ('Aimar Rizki','Jalan Angkasa no 27','11-11-2001','Online','pria'),
       ('Gilang Wahyu','Jalan Bumi no 13','07-08-1998','Online','pria'),
       ('Rafi Faiz','Jalan Planet no 65','03-10-2006','Online','pria'),
       ('Ami','Jalan Angkasa no 27','03-10-1978','Online','wanita');


# Membuat tabel product
create table product (
    id_product varchar(255) not null ,
    name_product varchar(255),
    price int not null ,
    primary key (id_product)
);

# insert into product
insert into product (id_product, name_product, price)
values ('P0001','Meja Belajar',150000),
       ('P0002','Kursi Belajar',125000),
       ('P0003','Lemari',350000);

select * from product;

alter table product
add column product_type varchar(255);


# Menambahkan tipe product ke kolom product


# Membuat tabel product
create table product_type (
    product_type_id varchar(255),
    type_product varchar(255),
    primary key (product_type_id)
);

select * from product_type;

alter table product
add constraint fk_product_type
foreign key (product_type) references product_type(product_type_id);

# Insert tabel product
insert into product_type (product_type_id, type_product)
values ('C0001','Perabotan Rumah Tangga'),
       ('C0002','Perhiasan Rumah'),
       ('C0003','Perabotan Rumah');

select * from product_type;

# insert category into product many-to-many
update product
set product_type = 'C0001'
where id_product = 'P0001';

update product
set product_type = 'C0001'
where id_product = 'P0002';

update product
set product_type = 'C0002'
where id_product = 'P0003';

select * from product
join product_type on (product.product_type = product_type.product_type_id);

# Membuat tabel operator
create table operators (
    id int not null auto_increment,
    value varchar(255),
    primary key (id)
);

# membuat tabel description
create table product_description (
    product_desctiption_id varchar(255) ,
    deskripsi varchar(255),
    primary key (product_desctiption_id)
);

# inser product description
insert into product_description(product_desctiption_id, deskripsi)
values ('info detail meja','ini meja belajar ukuran 60x40 cm'),
       ('info detail kursi','ini kursi belajar nyaman'),
       ('info detail lemari','ini lemari besar');

select * from product_description;


# Menambahkan column deskripsi ke tabel produk
alter table product
add column deskripsi_produk varchar(255);

select * from product;

alter table product
add constraint fk_product_description_id
foreign key (deskripsi_produk) references product_description(product_desctiption_id);

update  product
set deskripsi_produk='info detail meja'
where id_product = 'P0001';

update  product
set deskripsi_produk='info detail kursi'
where id_product = 'P0002';

update  product
set deskripsi_produk='info detail lemari'
where id_product = 'P0003';

# membuat tabel payment method
create table payment_method (
    id int not null auto_increment,
    payment_name varchar(255),
    biaya_layanan varchar(255),
    primary key (id)
);

insert into payment_method (payment_name, biaya_layanan)
VALUES ('Transfer Bank Manual','+1000'),
       ('Kartu Debit/Credit','+1500'),
       ('E-wallet','1000'),
       ('COD','2000');

select * from payment_method;

# membuat tabel transaction
create table transcantion (
    id_transaksi int not null auto_increment ,
    id_produk varchar(255) not null ,
    primary key (id_transaksi)
);

drop table transcantion;

alter table transcantion
add constraint fk_transaction_id
foreign key (id_produk) references product(id_product);

# Menambah data transaksi
insert into transcantion(id_produk)
values ('P0001');

insert into transcantion(id_produk)
values ('P0002');

select * from transcantion
JOIN product on (transcantion.id_produk = product.id_product);

# membuat tabel transaction_detail
create table transcantion_detail (
    id int not null auto_increment,
    id_produk varchar(255) not null ,
    id_price int not null ,
    primary key (id)
);


alter table transcantion_detail
add constraint fk_id
foreign key (id) references transcantion(id_transaksi);

alter table transcantion_detail
add constraint fk_id_produk
foreign key (id_produk) references product(id_product);

select * from transcantion_detail
JOIN transcantion on (transcantion_detail.id = transcantion.id_transaksi)
join product on (transcantion_detail.id_produk = product.id_product);

insert into transcantion_detail(id_produk, id_price)
values ('P0001',150000);



# membuat tabel kurir
create table kurir(
    id int not null,
    name varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);


# menambahkan kolum ongkos-dasar ke tabel kurir
alter table kurir
add column ongkos_dasar int not null;

# mengganti nama tabel kurir ke shippinh
rename table kurir to shipping;

# menghapus tombol kurir
drop table shipping;


# Soal nomor 7a
create table payment_method_description (
    id int not null auto_increment,
    payment_id int not null ,
    description_payment varchar(255),
    primary key (id),
    unique key payment_method_unique (payment_id),
    foreign key (payment_id) references payment_method(id)
);

select * from payment_method;

insert into payment_method_description (payment_id, description_payment)
values (1,'Pembayaran melalui transfer bank manual dengan biaya RP.1000');

insert into payment_method_description (payment_id, description_payment)
values (3,'Pembayaran praktis dengan ewallet biaya RP.1000');
insert into payment_method_description (payment_id, description_payment)
values (2,'Pembayaran simple cashless dengam biaya RP.1500');
insert into payment_method_description (payment_id, description_payment)
values (4,'Pembayaran di tempat  biaya RP.2000');

select * from payment_method_description
join payment_method on (payment_method_description.payment_id = payment_method.id);


# Soal nomor 7b
create table alamat(
    kode_pos int not null,
    kota varchar(255),
    primary key (kode_pos)
);

insert into alamat (kode_pos,kota)
values (14310,'jakarta'),
       (14320,'bogor'),
       (14330,'depok'),
       (14340,'tangerang'),
       (14350,'bekasi');

select * from alamat;

alter table user
add column kode_pos int  ;

select * from user;


alter table user
add constraint fk_kode_pos
foreign key (kode_pos) references alamat(kode_pos);

update user
set kode_pos = 14310
where id = 1;

update user
set kode_pos = 14320
where id = 2;

update user
set kode_pos = 14330
where id = 3;

update user
set kode_pos = 14310
where id = 4;

# Soal nomor 7c
create table user_payment_method_detail(
    id int not null auto_increment,
    user_id int not null ,
    payment_method int not null ,
    primary key (id)
);


alter table user_payment_method_detail
add constraint fk_user_id
foreign key (user_id) references user(id);

alter table user_payment_method_detail
add constraint fk_payment_method_detail
foreign key (payment_method) references payment_method(id);


select * from user;

insert into user_payment_method_detail (user_id, payment_method)
values (1,2);

insert into user_payment_method_detail (user_id, payment_method)
values (2,4);

select * from user_payment_method_detail
join user on (user_payment_method_detail.user_id = user.id)
join payment_method on (user_payment_method_detail.payment_method = payment_method.id)
join payment_method_description pmd on payment_method.id = pmd.payment_id;



