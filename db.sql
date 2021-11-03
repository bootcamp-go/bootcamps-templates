create table buyers
(
    id             integer
        constraint employees_pk
            primary key autoincrement,
    card_number_id varchar,
    first_name     varchar,
    last_name      varchar
);

create unique index buyers_id_uindex
    on buyers (id);

create table employees
(
    id             integer
        constraint employees_pk
            primary key autoincrement,
    card_number_id varchar,
    first_name     varchar,
    last_name      varchar,
    warehouse_id   int
);

create unique index employees_id_uindex
    on employees (id);

create table products
(
    id                               INTEGER not null
        primary key autoincrement,
    description                      TEXT,
    expiration_rate                  INTEGER,
    freezing_rate                    INTEGER,
    height                           INTEGER not null,
    lenght                           INTEGER not null,
    netweight                        INTEGER not null,
    product_code                     TEXT    not null,
    recommended_freezing_temperature INTEGER,
    width                            INTEGER not null,
    id_product_type                  INTEGER,
    id_seller                        INTEGER
);

create table sections
(
    id                  integer
        constraint sections_pk
            primary key autoincrement,
    section_number      int,
    current_temperature int,
    minimum_temperature int,
    current_capacity    int,
    minimum_capacity    int,
    maximum_capacity    int,
    warehouse_id        int,
    product_type_id     int
);

create unique index sections_id_uindex
    on sections (id);

create table sellers
(
    id           integer not null
        constraint sellers_pk
            primary key autoincrement,
    cid          int     not null,
    company_name varchar not null,
    address      varchar not null,
    telephone    varchar not null,
    locality_id  int     not null
);

create unique index sellers_cid_uindex
    on sellers (cid);

create unique index sellers_id_uindex
    on sellers (id);

create table users
(
    id       INTEGER not null
        primary key autoincrement,
    password TEXT    not null,
    username TEXT    not null
);

create table warehouses
(
    id                  integer
        constraint warehouses_pk
            primary key autoincrement,
    address             varchar,
    telephone           varchar,
    warehouse_code      varchar,
    minimun_capacity    int,
    minimun_temperature int,
    section_number      int
);

create unique index warehouses_id_uindex
    on warehouses (id);

