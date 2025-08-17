INSERT INTO users (id, balance)
SELECT 1, 100.00
WHERE NOT EXISTS (SELECT 1 FROM users WHERE id = 1);

INSERT INTO users (id, balance)
SELECT 2, 200.00
WHERE NOT EXISTS (SELECT 1 FROM users WHERE id = 2);

INSERT INTO users (id, balance)
SELECT 3, 300.00
WHERE NOT EXISTS (SELECT 1 FROM users WHERE id = 3);

SELECT setval(pg_get_serial_sequence('users', 'id'), GREATEST(MAX(id), 1))
FROM users;
