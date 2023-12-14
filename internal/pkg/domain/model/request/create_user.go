package request

type CreateNewUser struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
