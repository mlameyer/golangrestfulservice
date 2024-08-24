package dto

type AuthenticationRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
