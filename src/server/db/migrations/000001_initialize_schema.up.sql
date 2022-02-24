-- initialize our schema and desired tables

CREATE TABLE IF NOT EXISTS users(
  user_id VARCHAR(100) PRIMARY KEY,
  expo_token VARCHAR(100) NOT NULL
);
CREATE TABLE IF NOT EXISTS locations(
  location_id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  coords geometry(Point, 4326) NOT NULL
);
CREATE TABLE IF NOT EXISTS entry_log(
  entry_id SERIAL PRIMARY KEY,
  user_id VARCHAR(100) REFERENCES users (user_id),
  location_id INTEGER REFERENCES locations (location_id),
  entry_time TIMESTAMP,
  exit_time TIMESTAMP
);

-- add locations
INSERT INTO locations (name, coords)
VALUES
(
    'The Spire',
    ST_SetSRID(ST_MakePoint(53.349813, -6.260254), 4326)
),
(
    'DCU Nubar',
    ST_SetSRID(ST_MakePoint(53.385006, -6.258963), 4326)
),
(
    'The Academy',
    ST_SetSRID(ST_MakePoint(53.348001, -6.261989), 4326)
);
