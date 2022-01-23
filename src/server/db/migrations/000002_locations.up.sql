CREATE TABLE IF NOT EXISTS locations(
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  coords geometry(Point, 4326) NOT NULL
);
