CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,

    order_id INT NOT NULL,        -- к какому заказу относится
    product_id INT NOT NULL,      -- какой товар купили

    quantity INT NOT NULL,        -- сколько купили
    price NUMERIC(10,2) NOT NULL, -- цена за 1 шт НА МОМЕНТ ПОКУПКИ

    created_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE
);
