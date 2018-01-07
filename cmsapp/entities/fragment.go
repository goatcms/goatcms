package entities

// FragmentAllFields is a array contains list of all Fragment fields (except ID)
var FragmentAllFields = []string{"Content", "Lang", "Name"}

// FragmentMainFields is a array contains list of main Fragment fields (except ID)
var FragmentMainFields = []string{"Content", "Lang", "Name"}

// FragmentSystemFields is a array contains list of system Fragment fields (except ID)
var FragmentSystemFields = []string{}

// Fragment storages fragment data
type Fragment struct {
	ID      *int64  `json:"id"`
	Content *string `json:"content" db:"Content"`
	Lang    *string `json:"lang" db:"Lang"`
	Name    *string `json:"name" db:"Name"`
}
