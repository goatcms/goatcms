package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// TranslationRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type TranslationRows struct {
	*sql.Rows
}

func NewTranslationRows(rows *sql.Rows) *TranslationRows {
	return &TranslationRows{
		Rows: rows,
	}
}

func (rows *TranslationRows) InjectTo(dest *entities.Translation) (err error) {
	var columns []string
	if columns, err = rows.Rows.Columns(); err != nil {
		return err
	}
	values := make([]interface{}, len(columns))
	for i, name := range columns {
		switch name {
		case "Value":
			values[i] = &dest.Value
		case "Key":
			values[i] = &dest.Key
		default:
			return fmt.Errorf("TranslationRows.InjectTo unknow field %v", name)
		}
	}
	if err = rows.Rows.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (rows *TranslationRows) Get() (entity *entities.Translation, err error) {
	entity = &entities.Translation{}
	if err = rows.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
