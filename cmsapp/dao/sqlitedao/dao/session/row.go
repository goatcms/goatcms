package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// SessionRow is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the row
type SessionRow struct {
	*sql.Row
	columns []string
}

func NewSessionRow(row *sql.Row, columns []string) *SessionRow {
	return &SessionRow{
		Row:     row,
		columns: columns,
	}
}

func (row *SessionRow) Columns() (values []string, err error) {
	return row.columns, nil
}

func (row *SessionRow) InjectTo(dest *entities.Session) (err error) {
	values := make([]interface{}, len(row.columns))
	for i, name := range row.columns {
		switch name {
		case "ID":
			values[i] = &dest.ID
		case "Secret":
			values[i] = &dest.Secret
		case "UserID":
			values[i] = &dest.UserID
		default:
			return fmt.Errorf("SessionRow.InjectTo unknow field %v", name)
		}
	}
	if err = row.Row.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (row *SessionRow) Get() (entity *entities.Session, err error) {
	entity = &entities.Session{}
	if err = row.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}

func (row *SessionRow) GetValues() (values []interface{}, err error) {
	var columns []string
	if columns, err = row.Columns(); err != nil {
		return nil, err
	}
	values = make([]interface{}, len(columns))
	for i, _ := range values {
		var reference interface{}
		values[i] = &reference
	}
	if err = row.Row.Scan(values...); err != nil {
		return nil, err
	}
	return values, nil
}
