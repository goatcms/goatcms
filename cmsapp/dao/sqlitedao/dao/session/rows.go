package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// SessionRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type SessionRows struct {
	*sql.Rows
}

func NewSessionRows(rows *sql.Rows) *SessionRows {
	return &SessionRows{
		Rows: rows,
	}
}

func (rows *SessionRows) InjectTo(dest *entities.Session) (err error) {
	var columns []string
	if columns, err = rows.Rows.Columns(); err != nil {
		return err
	}
	values := make([]interface{}, len(columns))
	for i, name := range columns {
		switch name {
		case "ID":
			values[i] = &dest.ID
		case "Secret":
			values[i] = &dest.Secret
		case "UserID":
			values[i] = &dest.UserID
		default:
			return fmt.Errorf("SessionRows.InjectTo unknow field %v", name)
		}
	}
	if err = rows.Rows.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (rows *SessionRows) Get() (entity *entities.Session, err error) {
	entity = &entities.Session{}
	if err = rows.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}

func (rows *SessionRows) GetValues() (values []interface{}, err error) {
	var columns []string
	if columns, err = rows.Columns(); err != nil {
		return nil, err
	}
	values = make([]interface{}, len(columns))
	for i, _ := range values {
		var reference interface{}
		values[i] = &reference
	}
	if err = rows.Rows.Scan(values...); err != nil {
		return nil, err
	}
	return values, nil
}
