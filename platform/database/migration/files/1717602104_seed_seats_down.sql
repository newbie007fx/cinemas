DELETE FROM seats WHERE (theater_id = (SELECT id FROM theaters WHERE name = 'THeater 1') AND seat_row = 'A' AND seat_column IN (1, 2))
  OR (theater_id = (SELECT id FROM theaters WHERE name = 'Theater 1') AND seat_row = 'B' AND seat_column = 1)
  OR (theater_id = (SELECT id FROM theaters WHERE name = 'Theater 2') AND seat_row = 'A' AND seat_column = 1)
  OR (theater_id = (SELECT id FROM theaters WHERE name = 'Studio 1') AND seat_row = 'A' AND seat_column IN (1, 2));
