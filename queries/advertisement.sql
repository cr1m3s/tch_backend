-- name: CreateAdvertisement :one
INSERT INTO advertisements (
  title,
  provider,
  provider_id,
  attachment,
  experience,
  category,
  time,
  price,
  format,
  language,
  description,
  mobile_phone,
  email,
  telegram,
  created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: GetAdvertisementAll :many
SELECT * FROM advertisements;

-- name: GetAdvertisementByID :one
SELECT * FROM advertisements
WHERE id = $1;

-- name: GetAdvertisementByUsername :many
SELECT * FROM advertisements
WHERE provider = $1;

-- name: GetAdvertisementByUserID :many
SELECT * FROM advertisements
WHERE provider_id = $1;

-- name: GetAdvertisementByCategory :many
SELECT * FROM advertisements
WHERE category = $1;

SELECT * FROM advertisements
WHERE time <= $1;

SELECT * FROM advertisements
WHERE format = $1;

SELECT * FROM advertisements
WHERE experience >= $1
AND experience <= $2;

SELECT * FROM advertisements
WHERE language = $1;

-- name: UpdateAdvertisement :one
UPDATE advertisements
set   title = $2,
created_at = $3,
attachment = $4,
experience = $5,
category = $6,
time = $7,
price = $8,
format = $9,
language = $10,
description = $11,
mobile_phone = $12,
telegram = $13
WHERE id = $1
RETURNING *;

-- name: DeleteAdvertisementByID :exec
DELETE FROM advertisements
WHERE id = $1;

-- name: DeleteAdvertisementByUserID :exec
DELETE FROM advertisements
WHERE provider_id = $1;

-- name: FilterAdvertisements :many
SELECT * FROM advertisements
WHERE
    (category IS NULL OR (table.category = category))
    AND ($time IS NULL OR (table.time <= $tme))
    AND ($format IS NULL OR (table.format = $format))
    AND (($4 IS NULL AND $5 IS NULL) OR (experience >= $4 AND table.experience <= $5))
    AND ($language IS NULL OR (table.language = $language));