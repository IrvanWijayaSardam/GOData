package models

// Post type details
type Post struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Telp      string `json:"telp"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
