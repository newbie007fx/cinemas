INSERT INTO theaters (cinema_id, name)
VALUES
  ((SELECT id FROM cinemas WHERE name = 'Bioskop Semarang OK'), 'Theater 1'),
  ((SELECT id FROM cinemas WHERE name = 'Bioskop Semarang OK'), 'Theater 2'),
  ((SELECT id FROM cinemas WHERE name = 'Bioskop Solo Baru'), 'Studio 1');
