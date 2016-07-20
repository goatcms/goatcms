package models

import (
	"time"

	"github.com/goatcms/goatcms/models"
)

const (
	// ArticleDAOID is name used as article dao identifier
	ArticleDAOID = "articleDAO"
	// UserDAOID is name used as user dao identifier
	UserDAOID = "userDAO"
	// ImageDAOID is name user as image dao identifier
	ImageDAOID = "imageDAO"
)

// ArticleDTO represents an article entity
type ArticleDTO interface {
	GetID() int
	GetTitle() string
	GetContent() string
}

// ArticleDAO provide api to article access
type ArticleDAO interface {
	FindAll() []ArticleDTO
	FindByID(id int) ArticleDTO
	PersistAll(items []ArticleDTO)
}

// UserDTO represents a user entity
type UserDTO interface {
	GetID() int
	GetEmail() string
	GetPassHash() string
}

// UserDAO provide api to user access
type UserDAO interface {
	FindAll() []UserDTO
	FindByEmail(email string) UserDTO
	PersistAll(items []UserDTO)
}

// ImageDTO represents an image entity
type ImageDTO interface {
	GetID() int
	GetArticleID() int
	GetName() string
	GetLocation() string
	GetSize() int64
	GetCreatedAt() time.Time
	GetDescription() string
}

// ImageDAO provide api to image access
type ImageDAO interface {
	FindAll() []ImageDTO
	FindByID(id int) (models.ImageDTO, error)
	FindAllByArticleID(articleID int) []ImageDTO
	PersistOne(item models.ImageDTO) error
	PersistAll(items []ImageDTO)
}
