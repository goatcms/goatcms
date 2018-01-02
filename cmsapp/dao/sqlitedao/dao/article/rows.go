package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// ArticleRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type ArticleRows struct {
	*sql.Rows
}

func NewArticleRows(rows *sql.Rows) *ArticleRows {
	return &ArticleRows{
		Rows: rows,
	}
}

func (rows *ArticleRows) InjectTo(dest *entities.Article) (err error) {
	var columns []string
	if columns, err = rows.Rows.Columns(); err != nil {
		return err
	}
	values := make([]interface{}, len(columns))
	for i, name := range columns {
		switch name {
		case "Title":
			values[i] = &dest.Title
		case "Content":
			values[i] = &dest.Content
		default:
			return fmt.Errorf("ArticleRows.InjectTo unknow field %v", name)
		}
	}
	if err = rows.Rows.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (rows *ArticleRows) Get() (entity *entities.Article, err error) {
	entity = &entities.Article{}
	if err = rows.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
