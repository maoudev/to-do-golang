package domain

// DefaultCredentials represents the combination of email/password.
type DefaultCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
