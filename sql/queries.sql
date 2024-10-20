-- name: GetCatboxAuth :one
SELECT * FROM catbox_auth
WHERE ID = ? LIMIT 1;

-- name: CreateCatboxAuth :one
INSERT INTO catbox_auth (
  catbox_userhash
) VALUES ( 
    ?
)
RETURNING *;

-- name: UpdateCatboxAuth :exec
UPDATE catbox_auth
set catbox_userhash = ?
WHERE ID = ?
RETURNING *;

-- name: DeleteCatboxAuth :exec
DELETE FROM catbox_auth
WHERE id = ?;

