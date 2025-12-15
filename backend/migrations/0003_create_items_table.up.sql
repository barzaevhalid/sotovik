CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR,
    price INTEGER NOT NULL,           -- цена в рублях, целое число
    wholesale_price INTEGER,          -- оптовая цена
    is_active BOOLEAN DEFAULT true,
    in_stock BOOLEAN DEFAULT true,
    category_id INTEGER NOT NULL REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);