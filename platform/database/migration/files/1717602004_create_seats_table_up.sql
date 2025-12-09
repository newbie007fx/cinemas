CREATE TABLE IF NOT EXISTS seats (
  id SERIAL PRIMARY KEY,
  theater_id INT NOT NULL REFERENCES theaters(id) ON DELETE CASCADE,
  seat_row CHAR(1) NOT NULL,
  seat_column INT NOT NULL,
  CONSTRAINT unique_seat_in_theater UNIQUE (theater_id, seat_row, seat_column)
);
