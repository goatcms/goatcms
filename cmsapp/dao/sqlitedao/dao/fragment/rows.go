package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// FragmentRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type FragmentRows struct {
	*sql.Rows
}

func NewFragmentRows(rows *sql.Rows) *FragmentRows {
	return &FragmentRows{
		Rows: rows,
	}
}

func (rows *FragmentRows) InjectTo(dest *entities.Fragment) (err error) {
	var columns []string
	if columns, err = rows.Rows.Columns(); err != nil {
		return err
	}
	values := make([]interface{}, len(columns))
	for i, name := range columns {
		switch name {
		case "Lang":
			values[i] = &dest.Lang
		case "Name":
			values[i] = &dest.Name
		case "Content":
			values[i] = &dest.Content
		default:
			return fmt.Errorf("FragmentRows.InjectTo unknow field %v", name)
		}
	}
	if err = rows.Rows.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (rows *FragmentRows) Get() (entity *entities.Fragment, err error) {
	entity = &entities.Fragment{}
	if err = rows.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
