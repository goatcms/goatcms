package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// UserRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type UserRows struct {
	*sql.Rows
}

func NewUserRows(rows *sql.Rows) *UserRows {
	return &UserRows{
		Rows: rows,
	}
}

func (rows *UserRows) InjectTo(dest *entities.User) (err error) {
	var columns []string
	if columns, err = rows.Rows.Columns(); err != nil {
		return err
	}
	values := make([]interface{}, len(columns))
	for i, name := range columns {
		switch name {
		case "Email":
			values[i] = &dest.Email
		case "Password":
			values[i] = &dest.Password
		case "Firstname":
			values[i] = &dest.Firstname
		case "Login":
			values[i] = &dest.Login
		default:
			return fmt.Errorf("UserRows.InjectTo unknow field %v", name)
		}
	}
	if err = rows.Rows.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (rows *UserRows) Get() (entity *entities.User, err error) {
	entity = &entities.User{}
	if err = rows.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
