-- initialize our schema and desired tables

CREATE TABLE IF NOT EXISTS users(
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  phone_number VARCHAR(10) NOT NULL,
  covid_positive BOOLEAN NOT NULL
);
CREATE TABLE IF NOT EXISTS locations(
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  coords geometry(Point, 4326) NOT NULL
);
