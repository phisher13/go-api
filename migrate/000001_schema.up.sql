CREATE TABLE users (
    uuid uuid primary key default gen_random_uuid(),
    username text unique not null,
    email text unique not null,
    password_hash text not null
);

CREATE TABLE product (
    uuid uuid primary key default gen_random_uuid(),
    title text not null unique,
    description text,
    price int not null,
    created_at date default CURRENT_DATE,
    updated_at date default CURRENT_DATE,
    user_uuid uuid REFERENCES users(uuid)
);