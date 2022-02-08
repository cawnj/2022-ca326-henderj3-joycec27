-- insert some test data to ensure table validity

INSERT INTO users (name, phone_number, covid_positive)
VALUES
(
    'Conor Joyce',
    '0851234123',
    FALSE
),
(
    'Jason Henderson',
    '0854567456',
    FALSE
),
(
    'John Appleseed',
    '0857890789',
    FALSE
);
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
INSERT INTO entry_log (user_id, location_id, entry_time, exit_time)
VALUES
(
    1,
    3,
    DATE_TRUNC('second', NOW()::timestamp),
    DATE_TRUNC('second', (NOW() + interval '1 hour')::timestamp)
),
(
    2,
    2,
    DATE_TRUNC('second', NOW()::timestamp),
    DATE_TRUNC('second', (NOW() + interval '1 hour')::timestamp)
),
(
    3,
    2,
    DATE_TRUNC('second', (NOW() - interval '1 hour')::timestamp),
    DATE_TRUNC('second', (NOW() + interval '30 minutes')::timestamp)
);
