package authentication

type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type User struct {
	ID        int
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
