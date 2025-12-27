CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,          -- уникальный идентификатор строки заказа

    order_id INT NOT NULL,          -- к какому заказу относится
    product_id INT NOT NULL,        -- какой товар купили

    quantity INT NOT NULL CHECK (quantity > 0),  -- сколько купили
    price INT NOT NULL CHECK (price >= 0),      -- цена за 1 шт НА МОМЕНТ ПОКУПКИ (в рублях)

    created_at TIMESTAMP NOT NULL DEFAULT now(),

    -- связи
    CONSTRAINT fk_order_items_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_order_items_product
        FOREIGN KEY (product_id)
        REFERENCES items(id)
        ON DELETE RESTRICT,

    -- уникальность: один товар в одном заказе только один раз
    CONSTRAINT uniq_order_product
        UNIQUE (order_id, product_id)
);

-- индексы для ускорения запросов
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_product_id ON order_items(product_id);
