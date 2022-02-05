-- initialize our schema and desired tables

CREATE TABLE IF NOT EXISTS users(
  user_id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  phone_number VARCHAR(10) NOT NULL,
  covid_positive BOOLEAN NOT NULL
);
CREATE TABLE IF NOT EXISTS locations(
  location_id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  coords geometry(Point, 4326) NOT NULL
);
CREATE TABLE IF NOT EXISTS entry_log(
  entry_id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users (user_id),
  location_id INTEGER REFERENCES locations (location_id),
  entry_time TIMESTAMP,
  exit_time TIMESTAMP
);
