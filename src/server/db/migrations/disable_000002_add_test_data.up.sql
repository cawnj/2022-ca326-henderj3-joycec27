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
),
(
    'Jane Doe',
    '0871234123',
    FALSE
),
(
    'Joao Pereira',
    '0874567456',
    FALSE
),
(
    'Aaron Cleary',
    '0878907789',
    FALSE
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
(
    3,
    3,
    DATE_TRUNC('second', '2022-01-01 15:00:00'::timestamp),
    DATE_TRUNC('second', '2022-01-01 17:30:00'::timestamp)    

),
-- user 4, entry after, exit before
(
    4,
    3,
    DATE_TRUNC('second', '2022-01-01 16:30:00'::timestamp),
    DATE_TRUNC('second', '2022-01-01 16:45:00'::timestamp)    

),
-- user 5, entry after, exit after
(
    5,
    3,
    DATE_TRUNC('second', '2022-01-01 16:30:00'::timestamp),
    DATE_TRUNC('second', '2022-01-01 17:30:00'::timestamp)    

),
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
