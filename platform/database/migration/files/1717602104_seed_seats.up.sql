INSERT INTO seats (theater_id, seat_row, seat_column)
VALUES
  ((SELECT id FROM theaters WHERE name = 'Theater 1'), 'A', 1),
  ((SELECT id FROM theaters WHERE name = 'Theater 1'), 'A', 2),
  ((SELECT id FROM theaters WHERE name = 'Theater 1'), 'B', 1),
  ((SELECT id FROM theaters WHERE name = 'Theater 2'), 'A', 1),
  ((SELECT id FROM theaters WHERE name = 'Studio 1'), 'A', 1),
  ((SELECT id FROM theaters WHERE name = 'Studio 1'), 'A', 2);
