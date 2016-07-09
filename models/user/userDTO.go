package useremodel

// UserDTO is describing entity of user
type UserDTO struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	PassHash string `json:"passhash"`
}

// GetID return user's id
func (u *UserDTO) GetID() int {
	return u.ID
}

// GetEmail return user's email
func (u *UserDTO) GetEmail() string {
	return u.Email
}

// GetPassHash return user's hashed password
func (u *UserDTO) GetPassHash() string {
	return u.PassHash
}
