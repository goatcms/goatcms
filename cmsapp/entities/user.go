package entities

// UserMainFields is a array object contains list of all User fields without ID
var UserMainFields = []string{"Lastname", "Firstname"}

// User storages user data
type User struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname" db:"Firstname"`
	Lastname  string `json:"lastname" db:"Lastname"`
}
