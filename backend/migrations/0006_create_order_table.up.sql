CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    total INT DEFAULT 0,
    delivery_name VARCHAR(255),      -- имя получателя
    delivery_phone VARCHAR(50),      -- телефон получателя
    delivery_address TEXT,           -- адрес доставки
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
