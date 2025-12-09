INSERT INTO cinemas (city_id, name, address)
VALUES
  ((SELECT id FROM cities WHERE name = 'Semarang'), 'Bioskop Semarang OK', 'Jl. Yang Lurus'),
  ((SELECT id FROM cities WHERE name = 'Solo'), 'Bioskop Solo Baru', 'Jl. Lima puluh meter lagi');
