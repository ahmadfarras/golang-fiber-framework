package request

type CreateUser struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
