CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role_user TEXT NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now()
);
