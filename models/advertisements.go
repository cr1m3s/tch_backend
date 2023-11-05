package models

type AdvertisementInput struct {
	Title       string `json:"title"`
	Attachment  string `json:"attachment"`
	Experience  string `json:"experience"`
	Category    string `json:"category"`
	Time        string `json:"time"`
	Price       int32  `json:"price"`
	Format      string `json:"format"`
	Language    string `json:"language"`
	Description string `json:"description"`
	MobilePhone string `json:"mobile_phone"`
	Telegram    string `json:"telegram"`
}

type AdvertisementUpdate struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Attachment  string `json:"attachment"`
	Experience  string `json:"experience"`
	Category    string `json:"category"`
	Time        string `json:"time"`
	Price       int32  `json:"price"`
	Format      string `json:"format"`
	Language    string `json:"language"`
	Description string `json:"description"`
	MobilePhone string `json:"mobile_phone"`
	Telegram    string `json:"telegram"`
}

type AdvertisementDelete struct {
	AdvID int64 `json:"id"`
}

type AdvertisementFilter struct {
	Category string `json:"category"`
	Time     string `json:"time"`
	Format   string `json:"format"`
	MinExp   uint8  `json:"min_exp"`
	MaxExp   uint8  `json:"max_exp"`
	Language string `json:"language"`
}
