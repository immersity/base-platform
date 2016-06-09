create table account (
    id integer primary key auto_increment,
    role enum('user', 'admin') not null default 'user',
    verified boolean not null default false,
    email varchar(255) not null unique,
    password varchar(255) not null,
    first_name varchar(45) not null,
    last_name varchar(45) not null,
    created_on timestamp not null default current_timestamp,
    updated_on timestamp not null default current_timestamp on update current_timestamp
) engine=InnoDB;
