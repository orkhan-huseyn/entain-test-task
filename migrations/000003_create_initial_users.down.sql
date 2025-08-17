DELETE FROM users WHERE id IN (1, 2, 3);

SELECT setval(pg_get_serial_sequence('users', 'id'), COALESCE(MAX(id), 1))
FROM users;
