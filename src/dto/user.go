package dto

type InputLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type InputUserRegister struct {
	FullName  string `json:"full_name"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,gte=8"`
}