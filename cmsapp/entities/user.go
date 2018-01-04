package entities

// UserAllFields is a array contains list of all User fields (except ID)
var UserAllFields = []string{"Password", "Login", "Roles", "Lastname", "Firstname", "Email"}

// UserMainFields is a array contains list of main User fields (except ID)
var UserMainFields = []string{"Login", "Lastname", "Firstname", "Email"}

// UserSystemFields is a array contains list of system User fields (except ID)
var UserSystemFields = []string{"Roles", "Password"}

// User storages user data
type User struct {
	ID        *int64  `json:"id"`
	Roles     *string `json:"roles" db:"Roles"`
	Lastname  *string `json:"lastname" db:"Lastname"`
	Firstname *string `json:"firstname" db:"Firstname"`
	Email     *string `json:"email" db:"Email"`
	Password  *string `json:"password" db:"Password"`
	Login     *string `json:"login" db:"Login"`
}
