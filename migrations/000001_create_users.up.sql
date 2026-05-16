CREATE SCHEMA apiserver;

CREATE TABLE apiserver.users(
    id SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    encrypted_password VARCHAR(255) NOT NULL
);

CREATE TABLE apiserver.users_sessions(
    session_token VARCHAR(255) NOT NULL PRIMARY KEY,
    csrf_token VARCHAR(255) NOT NULL,
    user_id INT NOT NULL REFERENCES apiserver.users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,

    CONSTRAINT check_expires CHECK(expires_at > created_at)
);