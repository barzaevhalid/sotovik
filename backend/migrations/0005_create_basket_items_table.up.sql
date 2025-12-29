CREATE TABLE basket_items (
    id SERIAL PRIMARY KEY,
    basket_id INTEGER NOT NULL REFERENCES baskets(id),
    item_id INTEGER NOT NULL REFERENCES items(id),
    quantity INTEGER NOT NULL DEFAULT 1 CHECK (quantity > 0),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE UNIQUE INDEX idx_basket_item_unique ON basket_items(basket_id, item_id);
