package dao

import (
	"database/sql"
	"fmt"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

// FragmentRow is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the row
type FragmentRow struct {
	*sql.Row
	columns []string
}

func NewFragmentRow(row *sql.Row, columns []string) *FragmentRow {
	return &FragmentRow{
		Row:     row,
		columns: columns,
	}
}

func (row *FragmentRow) Columns() (values []string, err error) {
	return row.columns, nil
}

func (row *FragmentRow) InjectTo(dest *entities.Fragment) (err error) {
	values := make([]interface{}, len(row.columns))
	for i, name := range row.columns {
		switch name {
		case "ID":
			values[i] = &dest.ID
		case "Name":
			values[i] = &dest.Name
		case "Lang":
			values[i] = &dest.Lang
		case "Content":
			values[i] = &dest.Content
		default:
			return fmt.Errorf("FragmentRow.InjectTo unknow field %v", name)
		}
	}
	if err = row.Row.Scan(values...); err != nil {
		return err
	}
	return nil
}

func (row *FragmentRow) Get() (entity *entities.Fragment, err error) {
	entity = &entities.Fragment{}
	if err = row.InjectTo(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
