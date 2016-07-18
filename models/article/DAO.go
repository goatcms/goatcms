package articlemodel

import (
	"log"
	"strconv"

	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// ArticleDAO is describing entity of article
type ArticleDAO struct {
	db services.Database
}

// NewArticleDAO create new article DAO
func NewArticleDAO(db services.Database) models.ArticleDAO {
	return models.ArticleDAO(&ArticleDAO{
		db: db,
	})
}

// FindAll obtain all articles from database
func (dao *ArticleDAO) FindAll() []models.ArticleDTO {
	query := `
		SELECT id, title, content FROM articles
		`
	rows, err := dao.db.Adapter().Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []models.ArticleDTO
	for rows.Next() {
		item := ArticleDTO{}
		err2 := rows.Scan(&item.ID, &item.Title, &item.Content)
		if err2 != nil {
			log.Fatal(err2)
		}
		result = append(result, models.ArticleDTO(&item))
	}
	return result
}

// FindByID obtain article of given ID from database
func (dao *ArticleDAO) FindByID(id int) models.ArticleDTO {
	idString := strconv.FormatInt(int64(id), 10)
	query := `
		SELECT id, title, content FROM articles
		WHERE id = ?`
	rows, err := dao.db.Adapter().Query(query, idString)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result models.ArticleDTO
	for rows.Next() {
		item := ArticleDTO{}
		err2 := rows.Scan(&item.ID, &item.Title, &item.Content)
		if err2 != nil {
			log.Fatal(err2)
		}
		result = models.ArticleDTO(&item)
	}
	return result
}

// PersistAll store given articles to database
func (dao *ArticleDAO) PersistAll(items []models.ArticleDTO) {
	query := `
	INSERT OR REPLACE INTO articles(
		title, content
	) values(?, ?)
	`
	stmt, err := dao.db.Adapter().Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.GetTitle(), item.GetContent())
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
