package dto

type GitHubUserRequest struct {
	RedirectURI string `json:"redirect_uri"`
	Nonce       string `json:"nonce"`
	ExpiresAt   int64  `json:"expires_at"`
}
