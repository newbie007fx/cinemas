CREATE TABLE IF NOT EXISTS seat_inventories (
  id SERIAL PRIMARY KEY,
  showtime_id INT NOT NULL REFERENCES showtimes(id) ON DELETE CASCADE,
  seat_id INT NOT NULL REFERENCES seats(id) ON DELETE CASCADE,
  status VARCHAR(20) NOT NULL,
  version INT NOT NULL DEFAULT 1,
  reserved_until TIMESTAMP WITH TIME ZONE,
  CONSTRAINT unique_showtime_seat UNIQUE (showtime_id, seat_id)
);
