CREATE TABLE item_images (
    id SERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    image_url VARCHAR(255) NOT NULL,
    is_main BOOLEAN DEFAULT false,
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT now()
);
CREATE UNIQUE INDEX uniq_item_main_image
ON item_images (item_id)
WHERE is_main = true;
