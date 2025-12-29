CREATE TABLE items (
    id BIGSERIAL PRIMARY KEY,

    title VARCHAR(255) NOT NULL,
    description TEXT,

    price INTEGER NOT NULL CHECK (price >= 0),          -- цена в КОПЕЙКАХ
    wholesale_price INTEGER CHECK (wholesale_price >= 0),

    is_active BOOLEAN NOT NULL DEFAULT true,
    in_stock INTEGER NOT NULL DEFAULT 0 CHECK (in_stock >= 0),

    category_id BIGINT NOT NULL REFERENCES categories(id),

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);
