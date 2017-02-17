package models

type User struct {
	ID           int    `json:"id" db:"id" sqltype:"!int!primary!auto"`
	Email        string `json:"email" db:"Email" sqltype:"!char(400)"`
	PasswordHash string `json:"passwordHash" db:"PasswordHash" sqltype:"!char(400)"`
}

// Article is a entity represent single article
type Article struct {
	ID      int64  `json:"id" db:"id" sqltype:"!int!primary!auto" form:"?ID"`
	Title   string `json:"title" db:"Title" sqltype:"!char(400)" form:"Title"`
	Content string `json:"content" db:"Content" sqltype:"!char(400)" form:"Content"`
	Image   string `json:"image" db:"Image" sqltype:"!char(400)"`
}

func ArticleFactory() interface{} {
	return &Article{}
}

// Fragment storage single webpage fragment
type Fragment struct {
	ID    int64  `json:"id" db:"id" schema:"id" sqltype:"!int!primary!auto"`
	Key   string `json:"key" db:"Key" schema:"Key" sqltype:"!char(400)"`
	Value string `json:"value" db:"Value" schema:"Value" sqltype:"!char(400)"`
}
