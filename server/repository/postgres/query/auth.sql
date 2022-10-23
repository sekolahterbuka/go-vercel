-- name: Register :one

INSERT INTO
    auth (username, email, password)
VALUES ($1, $2, $3) RETURNING *;

-- name: Login :one

SELECT * FROM auth WHERE id = $1 LIMIT 1;