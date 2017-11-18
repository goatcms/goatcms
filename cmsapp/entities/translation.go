package entities

// Translation storages translation data
type Translation struct {
	Key   string `json:"key" db:"key"`
	Value string `json:"value" db:"value"`
}
