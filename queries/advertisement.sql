-- name: CreateAdvertisement :one
INSERT INTO advertisements (
  title,
  provider,
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
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
)
RETURNING *;

SELECT * FROM advertisements
WHERE provider = $1;

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

SELECT * FROM advertisements
WHERE
    ($1 IS NULL OR category = $1)
    AND ($2 IS NULL OR time <= $2)
    AND ($3 IS NULL OR format = $3)
    AND ($4 IS NULL OR (experience >= $4 AND experience <= $5))
    AND ($6 IS NULL OR language = $6);


UPDATE advertisements
set   title = $2,
provider = $3,
attachment = $4,
experience = $5,
category = $6,
time = $7,
price = $8,
format = $9,
language = $10,
description = $11,
mobile_phone = $12,
email = $13,
telegram = $14,
created_at = $15
WHERE id = $1
RETURNING *;

DELETE FROM advertisements
WHERE id = $1;

DELETE FROM advertisements
WHERE provider = $1;