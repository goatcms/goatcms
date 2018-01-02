package entities

// UserMainFields is a array object contains list of all User fields without ID
var UserMainFields = []string{"Email", "Login", "Firstname", "Password"}

// User storages user data
type User struct {
	ID        *int64  `json:"id"`
	Email     *string `json:"email" db:"Email"`
	Password  *string `json:"password" db:"Password"`
	Login     *string `json:"login" db:"Login"`
	Firstname *string `json:"firstname" db:"Firstname"`
}
