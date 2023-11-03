package models

import "time"

type AdvertisementInput struct {
	Title       string    `json:"title"`
	Attachment  string    `json:"attachment"`
	Experience  string    `json:"experience"`
	Category    string    `json:"category"`
	Time        string	  `json:"time"`
	Price       int32     `json:"price"`
	Format      string    `json:"format"`
	Language    string    `json:"language"`
	Description string    `json:"description"`
	MobilePhone string    `json:"mobile_phone"`
	Telegram    string    `json:"telegram"`
}

type AdvertisementUpdate struct{
	ID			int64	  `json:"id"`
	Title       string    `json:"title"`
	CreatedAt	time.Time `json:"created_at"`
	Attachment  string    `json:"attachment"`
	Experience  string    `json:"experience"`
	Category    string    `json:"category"`
	Time        string	  `json:"time"`
	Price       int32     `json:"price"`
	Format      string    `json:"format"`
	Language    string    `json:"language"`
	Description string    `json:"description"`
	MobilePhone string    `json:"mobile_phone"`
	Telegram    string    `json:"telegram"`
}