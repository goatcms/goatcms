package helpers

import (
	"database/sql"
)

// Rows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type Rows struct {
	*sql.Rows
}

func NewRows(rows *sql.Rows) *Rows {
	return &Rows{
		Rows: rows,
	}
}

func (rows *Rows) GetValues() (values []interface{}, err error) {
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
