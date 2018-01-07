package forms

// SigninAllFields is a array contains list of all Signin fields (except ID)
var SigninAllFields = []string{"Username", "Password"}

// SigninMainFields is a array contains list of main Signin fields (except ID)
var SigninMainFields = []string{"Username", "Password"}

// Signin storages signin data
type Signin struct {
	Username *string `json:"username" form:"username"`
	Password *string `json:"password" form:"password"`
}
