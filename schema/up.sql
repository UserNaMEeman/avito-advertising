CREATE DATABASE prod_DB;
CREATE TABLE products
(
id serial not null unique,
title varchar(200) not null unique,
description varchar(1000),
url1 varchar(255),
url2 varchar(255),
url3 varchar(255),
price money,
date_create timestamp default CURRENT_TIMESTAMP
);
