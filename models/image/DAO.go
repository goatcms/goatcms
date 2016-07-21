package imagemodel

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// ImageDAO is describing entity of image
type ImageDAO struct {
	db services.Database
}

// NewImageDAO create new ImageDAO
func NewImageDAO(db services.Database) (models.ImageDAO, error) {
	return models.ImageDAO(&ImageDAO{
		db: db,
	}), nil
}

// FindAll obtain all images from database
func (dao *ImageDAO) FindAll() ([]models.ImageDTO, error) {
	query := `
		SELECT id, article_id, name, location, description, size, created_at
		FROM images ORDER BY created_at DESC
		`
	rows, err := dao.db.Adapter().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []models.ImageDTO{}
	for rows.Next() {
		item := ImageDTO{}
		err := rows.Scan(
			&item.ID, &item.ArticleID, &item.Name, &item.Location,
			&item.Description, &item.Size, &item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, models.ImageDTO(&item))
	}
	return result, nil
}

// FindByID obtaint image of given ID from database
func (dao *ImageDAO) FindByID(id int) (models.ImageDTO, error) {
	query := `
	SELECT id, article_id, name, location, description, size, created_at
	FROM images
	WHERE id = ? LIMIT 1
	`
	row := dao.db.Adapter().QueryRow(query, id)
	var result models.ImageDTO
	item := ImageDTO{}
	err := row.Scan(
		&item.ID, &item.ArticleID, &item.Name,
		&item.Location, &item.Size, &item.CreatedAt,
	)
	switch {
	case err == sql.ErrNoRows:
		log.Println("No images with id", id)
	case err != nil:
		return nil, err
	default:
		result = models.ImageDTO(&item)
	}
	return result, err
}

// FindAllByArticleID obtain all images of given article's ID from database
func (dao *ImageDAO) FindAllByArticleID(articleID int) ([]models.ImageDTO, error) {
	// TODO try to change method arg to models.ArticleDTO not it's ID
	artIDString := strconv.FormatInt(int64(articleID), 10)
	query := `
		SELECT id, article_id, name, location, description, size, created_at
		FROM images
		WHERE article_id = ?
		`
	rows, err := dao.db.Adapter().Query(query, artIDString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.ImageDTO
	for rows.Next() {
		item := ImageDTO{}
		err := rows.Scan(
			&item.ID, &item.ArticleID, &item.Name, &item.Location,
			&item.Description, &item.Size, &item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, models.ImageDTO(&item))
	}
	return result, nil
}

// PersistOne store given image to database (with ID given!)
func (dao *ImageDAO) PersistOne(item models.ImageDTO) error {
	query := `
	INSERT OR REPLACE INTO images(
		article_id, name, location, description, size, created_at
	) values(?, ?, ?, ?, ?, ?)
	`
	_, err := dao.db.Adapter().Exec(
		query,
		item.GetArticleID(),
		item.GetName(),
		item.GetLocation(),
		item.GetDescription(),
		item.GetSize(),
		item.GetCreatedAt(),
	)
	return err
}

// PersistAll store given images to database
func (dao *ImageDAO) PersistAll(items []models.ImageDTO) {
	query := `
	INSERT OR REPLACE INTO images(
		article_id, name, location, description, size, created_at
	) values(?, ?, ?, ?, ?, ?)
	`
	stmt, err := dao.db.Adapter().Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(
			item.GetArticleID(),
			item.GetName(),
			item.GetLocation(),
			item.GetDescription(),
			item.GetSize(),
			item.GetCreatedAt(),
		)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
