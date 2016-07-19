package imagemodel

import (
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
func (dao *ImageDAO) FindAll() []models.ImageDTO {
	query := `
		SELECT id, article_id, name, location, description, size, created_at
		FROM images
		`
	rows, err := dao.db.Adapter().Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []models.ImageDTO
	for rows.Next() {
		item := ImageDTO{}
		err2 := rows.Scan(
			&item.ID, &item.ArticleID, &item.Name, &item.Location,
			&item.Description, &item.Size, &item.CreatedAt,
		)
		if err2 != nil {
			log.Fatal(err2)
		}
		result = append(result, models.ImageDTO(&item))
	}
	return result
}

// FindAllByArticleID obtain all images of given article from database
func (dao *ImageDAO) FindAllByArticleID(articleID int) []models.ImageDTO {
	artIDString := strconv.FormatInt(int64(articleID), 10)
	query := `
		SELECT id, article_id, name, location, description, size, created_at
		FROM images
		WHERE article_id = ?
		`
	rows, err := dao.db.Adapter().Query(query, artIDString)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []models.ImageDTO
	for rows.Next() {
		item := ImageDTO{}
		err2 := rows.Scan(
			&item.ID, &item.ArticleID, &item.Name, &item.Location,
			&item.Description, &item.Size, &item.CreatedAt,
		)
		if err2 != nil {
			log.Fatal(err2)
		}
		result = append(result, models.ImageDTO(&item))
	}
	return result
}

// PersistAll store given images to database`
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
