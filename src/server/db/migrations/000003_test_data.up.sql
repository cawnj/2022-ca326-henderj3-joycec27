INSERT INTO users (name, phone_number, covid_positive)
VALUES
(
    'John Doe',
    '0851234123',
    FALSE
);

INSERT INTO locations (name, coords)
VALUES
(
    'The Spire',
    ST_SetSRID(ST_MakePoint(53.3498133103617, -6.260254521788297), 4326)
);
