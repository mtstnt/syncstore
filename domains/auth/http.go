package auth

type LoginRequest struct {
	Username    string `json:"username"`
	PasswordB64 string `json:"password"`
}
