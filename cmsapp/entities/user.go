package entities

// UserAllFields is a array contains list of all User fields (except ID)
var UserAllFields = []string{"Firstname", "Lastname", "Email", "Password", "Roles", "Username"}

// UserMainFields is a array contains list of main User fields (except ID)
var UserMainFields = []string{"Firstname", "Lastname", "Email", "Username", "ID", "ID", "ID", "ID"}

// UserSystemFields is a array contains list of system User fields (except ID)
var UserSystemFields = []string{"Password", "Roles"}

// User storages user data
type User struct {
	ID        *int64  `json:"id"`
	Firstname *string `json:"firstname" db:"Firstname"`
	Lastname  *string `json:"lastname" db:"Lastname"`
	Email     *string `json:"email" db:"Email"`
	Password  *string `json:"password" db:"Password"`
	Roles     *string `json:"roles" db:"Roles"`
	Username  *string `json:"username" db:"Username"`
}
