package entities

// UserMainFields is a array object contains list of all User fields without ID
var UserMainFields = []string{"Login", "Firstname", "Password", "Email"}

// User storages user data
type User struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname" db:"Firstname"`
	Login     string `json:"login" db:"Login"`
	Email     string `json:"email" db:"Email"`
	Password  string `json:"password" db:"Password"`
}
