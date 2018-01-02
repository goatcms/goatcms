package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// ArticleRow is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the row
type ArticleRow struct {
	*sql.Row
	columns []string
}

func NewArticleRow(row *sql.Row, columns []string) *ArticleRow {
	return &ArticleRow{
		Row:     row,
		columns: columns,
	}
}

func (row *ArticleRow) Columns() (values []string, err error) {
	return row.columns, nil
}

func (row *ArticleRow) InjectTo(dest *entities.Article) (err error) {
	values := make([]interface{}, len(row.columns))
	for i, name := range row.columns {
		switch name {
		case "Content":
			values[i] = &dest.Content
		case "Title":
			values[i] = &dest.Title
		default:
			return fmt.Errorf("ArticleRow.InjectTo unknow field %v", name)
		}
	}
	if err = row.Row.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (row *ArticleRow) Get() (entity *entities.Article, err error) {
	entity = &entities.Article{}
	if err = row.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
