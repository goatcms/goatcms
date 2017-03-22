package models

type User struct {
	ID           int64  `json:"id" db:"id" sqltype:"!id" form:"?ID"`
	Email        string `json:"email" db:"Email" sqltype:"!char(400)" form:"Email"`
	PasswordHash string `json:"passwordHash" db:"PasswordHash" sqltype:"!char(400)"`
}

// Article is a entity represent single article
type Article struct {
	ID      int64  `json:"id" db:"id" sqltype:"!id" form:"?ID"`
	Title   string `json:"title" db:"title" sqltype:"!char(400)" form:"Title"`
	Content string `json:"content" db:"content" sqltype:"!char(400)" form:"Content"`
	Image   string `json:"image" db:"image" sqltype:"!char(400)"`
}

func ArticleFactory() interface{} {
	return &Article{}
}

// Fragment storage single webpage fragment
type Fragment struct {
	ID    int64  `json:"id" db:"id" schema:"id" sqltype:"!id"`
	Key   string `json:"key" db:"Key" schema:"Key" sqltype:"!char(400)"`
	Value string `json:"value" db:"Value" schema:"Value" sqltype:"!char(400)"`
}
