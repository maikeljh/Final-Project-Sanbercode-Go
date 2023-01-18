-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    category_id BIGINT,
    price BIGINT,
    description VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE cart (
    id SERIAL PRIMARY KEY,
    product_id BIGINT,
    count BIGINT,
    user_id BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE order (
    id SERIAL PRIMARY KEY,
    cart_id BIGINT,
    user_id BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    address VARCHAR(256),
    phone_number VARCHAR(256),
    username VARCHAR(256),
    password VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
-- +migrate StatementEnd