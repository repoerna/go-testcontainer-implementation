create table if not exists public.staff
(
    id    serial primary key,
    first_name  varchar(45)                                               not null,
    last_name   varchar(45)                                               not null,
    address_id  text                                                      not null,
    email       varchar(50) not null,
    active      boolean   default true                                    not null,
    username    varchar(16)                                               not null,
    password    varchar(40),
    last_update timestamp default now()                                   not null,
    picture     bytea
);

-- alter table public.staff
--     owner to postgres;


