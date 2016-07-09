package useremodel

import (
	"log"

	"github.com/s3c0nDD/goatcms/models"
	"github.com/s3c0nDD/goatcms/services"
)

// UserDAO is describing entity of user
type UserDAO struct {
	db services.Database
}

// NewUserDAO create new article DAO
func NewUserDAO(db services.Database) (models.UserDAO, error) {
	return models.UserDAO(&UserDAO{
		db: db,
	}), nil
}

// GetAll obtains all users from database
func (dao *UserDAO) GetAll() []models.UserDTO {
	query := `
		SELECT id, email, pass_hash FROM users
		`
	rows, err := dao.db.Adapter().Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []models.UserDTO
	for rows.Next() {
		item := UserDTO{}
		err2 := rows.Scan(&item.ID, &item.Email, &item.PassHash)
		if err2 != nil {
			log.Fatal(err2)
		}
		result = append(result, models.UserDTO(&item))
	}
	return result
}

// PersistAll store given users to database
func (dao *UserDAO) PersistAll(items []models.UserDTO) {
	query := `
	INSERT OR REPLACE INTO users(
		email, pass_hash
	) values(?, ?)
	`
	stmt, err := dao.db.Adapter().Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, item := range items {
		// TODO store hashed passwords !
		_, err2 := stmt.Exec(item.GetEmail(), item.GetPassHash())
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
