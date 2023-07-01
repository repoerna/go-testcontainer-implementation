create table if not exists public.users
(
    id       text not null
        primary key,
    email    text,
    password text not null
);