package models

type LoginRequest struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type SignUpRequest struct {
	FirstName   string `json:"first_name" required:"true"`
	LastName    string `json:"last_name" required:"true"`
	Email       string `json:"email" required:"true"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" required:"true"`
}
