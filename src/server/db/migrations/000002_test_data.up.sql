-- test data
INSERT INTO users (user_id, expo_token)
VALUES
(
    'test_user',
    'test_token'
);
INSERT INTO entry_log (user_id, location_id, entry_time, exit_time)
VALUES
(
    'test_user',
    2,
    DATE_TRUNC('second', '2022-01-01 16:00:00'::timestamp),
    DATE_TRUNC('second', '2022-01-01 17:00:00'::timestamp)
);
