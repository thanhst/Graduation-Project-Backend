package dto

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Status       string `json:"status"`
	Message      string `json:"message,omitempty"`
}
