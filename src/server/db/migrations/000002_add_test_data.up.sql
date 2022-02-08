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
-- user 1, covid event
(
    1,
    3,
    DATE_TRUNC('second', '2022-01-01 16:00:00'::timestamp),
    DATE_TRUNC('second', '2022-01-01 17:00:00'::timestamp)
),
-- user 2, entry before, exit before
(
    2,
    3,
    DATE_TRUNC('second', '2022-01-01 15:30:00'::timestamp),
    DATE_TRUNC('second', '2022-01-01 16:30:00'::timestamp)
),
-- user 3, entry before, exit after
-- user 4, entry after, exit before
-- user 5, entry after, exit after
-- user 6, unrelated before
(
    6,
    3,
    DATE_TRUNC('second', '2021-12-31 16:00:00'::timestamp),
    DATE_TRUNC('second', '2021-12-31 17:00:00'::timestamp)
),
-- user 7, unrelated after
(
    6,
    3,
    DATE_TRUNC('second', '2022-01-02 16:00:00'::timestamp),
    DATE_TRUNC('second', '2022-01-02 17:00:00'::timestamp)
);
