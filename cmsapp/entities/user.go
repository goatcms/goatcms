package entities

// UserAllFields is a array contains list of all User fields (except ID)
var UserAllFields = []string{"Password", "Lastname", "Username", "Firstname", "Roles", "Email"}

// UserMainFields is a array contains list of main User fields (except ID)
var UserMainFields = []string{"Lastname", "Username", "Firstname", "Email"}

// UserSystemFields is a array contains list of system User fields (except ID)
var UserSystemFields = []string{"Password", "Roles"}

// User storages user data
type User struct {
	ID        *int64  `json:"id"`
	Password  *string `json:"password" db:"Password"`
	Lastname  *string `json:"lastname" db:"Lastname"`
	Username  *string `json:"username" db:"Username"`
	Firstname *string `json:"firstname" db:"Firstname"`
	Roles     *string `json:"roles" db:"Roles"`
	Email     *string `json:"email" db:"Email"`
}
