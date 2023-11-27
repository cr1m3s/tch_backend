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
SELECT * FROM advertisements
ORDER BY created_at DESC
LIMIT 10;

-- name: GetAdvertisementMy :many
SELECT * FROM advertisements 
WHERE provider_id = $1;

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
WITH filtered_ads AS (
SELECT * FROM advertisements
  WHERE
        (NULLIF(sqlc.arg(advCategory)::text, '')::text IS NULL OR category = sqlc.arg(advCategory)::text)
        AND (NULLIF(sqlc.arg(timeLength)::int, 0) IS NULL OR time <= sqlc.arg(timeLength)::int)
        AND (NULLIF(sqlc.arg(advFformat)::text, '') IS NULL OR format = sqlc.arg(advFormat)::text)
        AND ((NULLIF(sqlc.arg(minExp)::int, 0) IS NULL AND NULLIF(sqlc.arg(maxExp)::int, 0) IS NULL) OR (experience >= sqlc.arg(minExp)::int AND experience <= sqlc.arg(maxExp)::int))
        AND ((NULLIF(sqlc.arg(minPrice)::int, 0) IS NULL AND NULLIF(sqlc.arg(maxPrice)::int, 0) IS NULL) OR (price >= sqlc.arg(minPrice)::int AND price <= sqlc.arg(maxPrice)::int))
        AND (NULLIF(sqlc.arg(advLanguage)::text, '') IS NULL OR language = sqlc.arg(advLanguage)::text)
        AND (NULLIF(sqlc.arg(titleKeyword)::text, '') IS NULL OR title ILIKE '%' || sqlc.arg(titleKeyword)::text || '%')
)
SELECT *,
    COUNT(*) OVER () AS total_items
FROM filtered_ads
ORDER BY
  ( CASE
    WHEN sqlc.arg(orderBy)::text = 'price' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(price AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'experience' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(experience AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'date' AND sqlc.arg(sortOrder)::text = 'desc' THEN CAST(created_at AS TEXT) END) DESC,
  ( CASE
    WHEN sqlc.arg(orderBy)::text = 'price' THEN CAST(price AS TEXT)
    WHEN sqlc.arg(orderBy)::text = 'experience' THEN CAST(experience AS TEXT)  
    ELSE CAST(created_at AS TEXT) END) ASC                                     
LIMIT sqlc.arg(limitAdv)::integer    
OFFSET sqlc.arg(offsetAdv)::integer; 