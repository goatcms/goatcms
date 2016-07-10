package models

const (
	// ArticleDAOID is name used as article dao identifier
	ArticleDAOID = "articleDAO"
	// UserDAOID is name used as user dao identifier
	UserDAOID = "userDAO"
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
	PersistAll(items []UserDTO)
}
