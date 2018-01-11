package entities

// SessionAllFields is a array contains list of all Session fields (except ID)
var SessionAllFields = []string{"Secret", "UserID"}

// SessionMainFields is a array contains list of main Session fields (except ID)
var SessionMainFields = []string{"Secret", "UserID"}

// SessionSystemFields is a array contains list of system Session fields (except ID)
var SessionSystemFields = []string{}

// Session storages session data
type Session struct {
	ID     *int64  `json:"id"`
	Secret *string `json:"secret" db:"Secret"`
	UserID *int64
	User   *User
}
