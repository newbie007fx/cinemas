CREATE TABLE IF NOT EXISTS order_items (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  seat_inventory_id INT NOT NULL REFERENCES seat_inventories(id) ON DELETE CASCADE,
  price NUMERIC(10, 2) NOT NULL
);
