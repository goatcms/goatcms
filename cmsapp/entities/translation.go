package entities

// TranslationMainFields is a array object contains list of all Translation fields without ID
var TranslationMainFields = []string{"Key", "Value"}

// Translation storages translation data
type Translation struct {
	ID    *int64  `json:"id"`
	Value *string `json:"value" db:"Value"`
	Key   *string `json:"key" db:"Key"`
}
