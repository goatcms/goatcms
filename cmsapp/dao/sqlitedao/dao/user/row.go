package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// UserRow is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the row
type UserRow struct {
	*sql.Row
	columns []string
}

func NewUserRow(row *sql.Row, columns []string) *UserRow {
	return &UserRow{
		Row:     row,
		columns: columns,
	}
}

func (row *UserRow) Columns() (values []string, err error) {
	return row.columns, nil
}

func (row *UserRow) InjectTo(dest *entities.User) (err error) {
	values := make([]interface{}, len(row.columns))
	for i, name := range row.columns {
		switch name {
		case "ID":
			values[i] = &dest.ID
		case "Email":
			values[i] = &dest.Email
		case "Lastname":
			values[i] = &dest.Lastname
		case "Roles":
			values[i] = &dest.Roles
		case "Firstname":
			values[i] = &dest.Firstname
		case "Username":
			values[i] = &dest.Username
		case "Password":
			values[i] = &dest.Password
		default:
			return fmt.Errorf("UserRow.InjectTo unknow field %v", name)
		}
	}
	if err = row.Row.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (row *UserRow) Get() (entity *entities.User, err error) {
	entity = &entities.User{}
	if err = row.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
