package articlemodel

// ArticleDTO is describing entity of article
type ArticleDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GetID return article's id
func (a *ArticleDTO) GetID() int {
	return a.ID
}

// GetTitle return article's title
func (a *ArticleDTO) GetTitle() string {
	return a.Title
}

// GetContent return article's HTML content code
func (a *ArticleDTO) GetContent() string {
	return a.Content
}
