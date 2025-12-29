CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'customer'
    CHECK (role IN ('customer', 'wholesaler', 'admin')),
    is_blocked BOOLEAN DEFAULT false,
    store VARCHAR(255),
    phone VARCHAR(50),
    
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);