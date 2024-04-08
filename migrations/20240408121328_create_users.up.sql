CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    email VARCHAR NOT NULL unique,
    Name VARCHAR NOT NULL,
    encrypted_password VARCHAR NOT NULL
);