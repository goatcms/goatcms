package entities

// Article storages article data
type Article struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}
