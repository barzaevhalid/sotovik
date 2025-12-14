CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    role VARCHAR DEFAULT 'customer',
    is_blocked BOOLEAN DEFAULT false,
    store VARCHAR,
    phone VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);