CREATE SCHEMA apiserver;

CREATE TABLE apiserver.users(
    id SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL
);