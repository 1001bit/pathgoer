CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(31) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    reg_date DATE DEFAULT now()
);