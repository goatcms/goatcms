package imagemodel

import "time"

// ImageDTO is describing entity of image
type ImageDTO struct {
	ID          int       `json:"id"`
	ArticleID   int       `json:"articleid"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"createdat"`
}

// GetID returns image's id
func (i *ImageDTO) GetID() int {
	return i.ID
}

// GetArticleID returns article's id, to which image belongs
func (i *ImageDTO) GetArticleID() int {
	return i.ArticleID
}

// GetName returns image's name
func (i *ImageDTO) GetName() string {
	return i.Name
}

// GetLocation returns image's location path
func (i *ImageDTO) GetLocation() string {
	return i.Location
}

// GetDescription returns image's description
func (i *ImageDTO) GetDescription() string {
	return i.Description
}

// GetSize returns image's size
func (i *ImageDTO) GetSize() int64 {
	return i.Size
}

// GetCreatedAt returns image's creation time
func (i *ImageDTO) GetCreatedAt() time.Time {
	return i.CreatedAt
}
