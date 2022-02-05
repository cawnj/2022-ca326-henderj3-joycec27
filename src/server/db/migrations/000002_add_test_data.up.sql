-- insert some test data to ensure table validity

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
INSERT INTO entry_log (user_id, location_id, entry_time, exit_time)
VALUES
(
    1,
    1,
    DATE_TRUNC('second', NOW()::timestamp),
    DATE_TRUNC('second', (NOW() + interval '1 hour')::timestamp)
);
