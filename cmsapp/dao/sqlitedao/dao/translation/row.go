package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// TranslationRow is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the row
type TranslationRow struct {
	*sql.Row
	columns []string
}

func NewTranslationRow(row *sql.Row, columns []string) *TranslationRow {
	return &TranslationRow{
		Row:     row,
		columns: columns,
	}
}

func (row *TranslationRow) Columns() (values []string, err error) {
	return row.columns, nil
}

func (row *TranslationRow) InjectTo(dest *entities.Translation) (err error) {
	values := make([]interface{}, len(row.columns))
	for i, name := range row.columns {
		switch name {
		case "Key":
			values[i] = &dest.Key
		case "Value":
			values[i] = &dest.Value
		default:
			return fmt.Errorf("TranslationRow.InjectTo unknow field %v", name)
		}
	}
	if err = row.Row.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (row *TranslationRow) Get() (entity *entities.Translation, err error) {
	entity = &entities.Translation{}
	if err = row.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
