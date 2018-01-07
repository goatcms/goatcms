package forms

// SignupAllFields is a array contains list of all Signup fields (except ID)
var SignupAllFields = []string{"Password", "Username", "Email", "Lastname", "Firstname"}

// SignupMainFields is a array contains list of main Signup fields (except ID)
var SignupMainFields = []string{"Password", "Username", "Email", "Lastname", "Firstname"}

// Signup storages signup data
type Signup struct {
	Email     *string         `json:"email" form:"email"`
	Lastname  *string         `json:"lastname" form:"lastname"`
	Password  *RepeatPassword `json:"password" form:"password"`
	Firstname *string         `json:"firstname" form:"firstname"`
	Username  *string         `json:"username" form:"username"`
}
