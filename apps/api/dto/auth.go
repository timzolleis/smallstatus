package dto

type CreateUserDTO struct {
	Name     string `json:"name" ,validate:"required"`
	Email    string `json:"email" ,validate:"required,email"`
	Password string `json:"password" ,validate:"required,min=6,max=32"`
}

type LoginDTO struct {
	Email    string `json:"email" ,validate:"required,email"`
	Password string `json:"password" ,validate:"required,min=6,max=32"`
}
