// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: advertisement.sql

package queries

import (
	"context"
	"time"
)

const createAdvertisement = `-- name: CreateAdvertisement :one
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
RETURNING id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at
`

type CreateAdvertisementParams struct {
	Title       string    `json:"title"`
	Provider    string    `json:"provider"`
	ProviderID  int64     `json:"provider_id"`
	Attachment  string    `json:"attachment"`
	Experience  string    `json:"experience"`
	Category    string    `json:"category"`
	Time        string    `json:"time"`
	Price       int32     `json:"price"`
	Format      string    `json:"format"`
	Language    string    `json:"language"`
	Description string    `json:"description"`
	MobilePhone string    `json:"mobile_phone"`
	Email       string    `json:"email"`
	Telegram    string    `json:"telegram"`
	CreatedAt   time.Time `json:"created_at"`
}

func (q *Queries) CreateAdvertisement(ctx context.Context, arg CreateAdvertisementParams) (Advertisement, error) {
	row := q.db.QueryRow(ctx, createAdvertisement,
		arg.Title,
		arg.Provider,
		arg.ProviderID,
		arg.Attachment,
		arg.Experience,
		arg.Category,
		arg.Time,
		arg.Price,
		arg.Format,
		arg.Language,
		arg.Description,
		arg.MobilePhone,
		arg.Email,
		arg.Telegram,
		arg.CreatedAt,
	)
	var i Advertisement
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Provider,
		&i.ProviderID,
		&i.Attachment,
		&i.Experience,
		&i.Category,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAdvertisementByID = `-- name: DeleteAdvertisementByID :exec
DELETE FROM advertisements
WHERE id = $1
`

func (q *Queries) DeleteAdvertisementByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAdvertisementByID, id)
	return err
}

const deleteAdvertisementByUserID = `-- name: DeleteAdvertisementByUserID :exec
DELETE FROM advertisements
WHERE provider_id = $1
`

func (q *Queries) DeleteAdvertisementByUserID(ctx context.Context, providerID int64) error {
	_, err := q.db.Exec(ctx, deleteAdvertisementByUserID, providerID)
	return err
}

const filterAdvertisements = `-- name: FilterAdvertisements :many
 SELECT id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
        WHERE
        (NULLIF($1, '')::text IS NULL OR category = $1::text)
        AND (NULLIF($2::text, '') IS NULL OR time <= $1::text)
        AND (NULLIF($3::text, '') IS NULL OR format = $3::text)
        AND ((NULLIF($4::text, '') IS NULL AND NULLIF($5::text, '') IS NULL) OR (experience >= $4::text AND experience <= $5::text))
        AND (NULLIF($6::text, '') IS NULL OR language = $6::text)
`

type FilterAdvertisementsParams struct {
	Category interface{} `json:"category"`
	Time     string      `json:"time"`
	Format   string      `json:"format"`
	Minexp   string      `json:"minexp"`
	Maxexp   string      `json:"maxexp"`
	Language string      `json:"language"`
}

func (q *Queries) FilterAdvertisements(ctx context.Context, arg FilterAdvertisementsParams) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, filterAdvertisements,
		arg.Category,
		arg.Time,
		arg.Format,
		arg.Minexp,
		arg.Maxexp,
		arg.Language,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.Category,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementAll = `-- name: GetAdvertisementAll :many
SELECT id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
`

func (q *Queries) GetAdvertisementAll(ctx context.Context) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.Category,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByCategory = `-- name: GetAdvertisementByCategory :many
SELECT id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE category = $1
`

func (q *Queries) GetAdvertisementByCategory(ctx context.Context, category string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByCategory, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.Category,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByID = `-- name: GetAdvertisementByID :one
SELECT id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE id = $1
`

func (q *Queries) GetAdvertisementByID(ctx context.Context, id int64) (Advertisement, error) {
	row := q.db.QueryRow(ctx, getAdvertisementByID, id)
	var i Advertisement
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Provider,
		&i.ProviderID,
		&i.Attachment,
		&i.Experience,
		&i.Category,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
	)
	return i, err
}

const getAdvertisementByUserID = `-- name: GetAdvertisementByUserID :many
SELECT id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE provider_id = $1
`

func (q *Queries) GetAdvertisementByUserID(ctx context.Context, providerID int64) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByUserID, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.Category,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByUsername = `-- name: GetAdvertisementByUsername :many
SELECT id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE provider = $1
`

func (q *Queries) GetAdvertisementByUsername(ctx context.Context, provider string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByUsername, provider)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.Category,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdvertisement = `-- name: UpdateAdvertisement :one
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
RETURNING id, title, provider, provider_id, attachment, experience, category, time, price, format, language, description, mobile_phone, email, telegram, created_at
`

type UpdateAdvertisementParams struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	Attachment  string    `json:"attachment"`
	Experience  string    `json:"experience"`
	Category    string    `json:"category"`
	Time        string    `json:"time"`
	Price       int32     `json:"price"`
	Format      string    `json:"format"`
	Language    string    `json:"language"`
	Description string    `json:"description"`
	MobilePhone string    `json:"mobile_phone"`
	Telegram    string    `json:"telegram"`
}

func (q *Queries) UpdateAdvertisement(ctx context.Context, arg UpdateAdvertisementParams) (Advertisement, error) {
	row := q.db.QueryRow(ctx, updateAdvertisement,
		arg.ID,
		arg.Title,
		arg.CreatedAt,
		arg.Attachment,
		arg.Experience,
		arg.Category,
		arg.Time,
		arg.Price,
		arg.Format,
		arg.Language,
		arg.Description,
		arg.MobilePhone,
		arg.Telegram,
	)
	var i Advertisement
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Provider,
		&i.ProviderID,
		&i.Attachment,
		&i.Experience,
		&i.Category,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
	)
	return i, err
}
