package models

import (
	"time"

	"github.com/goatcms/goat-core/db"
	"github.com/jmoiron/sqlx"
)

const (
	// ArticleDAOID is name used as article dao identifier
	ArticleDAOID = "dao.article"
	// UserDAOID is name used as user dao identifier
	UserDAOID = "userDAO"
	// ImageDAOID is name user as image dao identifier
	ImageDAOID = "imageDAO"
)

// ArticleEntity is a entity represent single article
type ArticleEntity struct {
	ID      int64  `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Image   string `json:"image" db:"image"`
}

// ArticleDAO provide api to article access
type ArticleDAO interface {
	db.DAO

	ToEntities(rows *sqlx.Rows) ([]*ArticleEntity, error)
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
	GetDescription() string
	GetSize() int64
	GetCreatedAt() time.Time
}

// ImageDAO provide api to image access
type ImageDAO interface {
	FindAll() ([]ImageDTO, error)
	FindByID(id int) (ImageDTO, error)
	FindAllByArticleID(articleID int) ([]ImageDTO, error)
	PersistOne(item ImageDTO) error
	PersistAll(items []ImageDTO)
}
