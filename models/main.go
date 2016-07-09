package models

const (
	// ArticleDAOID is name used as article dao identyfier
	ArticleDAOID = "articleDAO"
	// UserDAOID is name used as users dao identyfier
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
	GetAll() []ArticleDTO
	GetOne(id int) ArticleDTO
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
	GetAll() []UserDTO
	PersistAll(items []UserDTO)
}
