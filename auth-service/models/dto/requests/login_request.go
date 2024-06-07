package requests

/**
 * created by	: albym
 * created at	: 5 Jun 2024
 *
 * login request type
 */
type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=64"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}
