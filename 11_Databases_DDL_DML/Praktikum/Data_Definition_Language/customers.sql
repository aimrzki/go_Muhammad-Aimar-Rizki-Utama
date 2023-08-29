create table customers(
    id int not null auto_increment,
    email varchar(100) not null ,
    first_name varchar(100) not null ,
    last_name varchar(100),
    primary key (id),
    unique key email_unique (email)
);

describe customers;

alter table customers
drop constraint email_unique;

alter table customers
add constraint email_unique unique (email);

insert into customers(email, first_name, last_name)
values ('muhammadaimar77@gmail.com','Aimar','Rizki');

insert into customers(email, first_name, last_name)
values ('aimrzki@student.telkomuniversity.ac.id','Aimar','Rizki');

select * from customers;
