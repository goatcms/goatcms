package entities

// ArticleMainFields is a array object contains list of all Article fields without ID
var ArticleMainFields = []string{"Title", "Content"}

// Article storages article data
type Article struct {
	ID      *int64  `json:"id"`
	Content *string `json:"content" db:"Content"`
	Title   *string `json:"title" db:"Title"`
}
