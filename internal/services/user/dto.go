package user

type CreateUserInput struct {
	Email    string `json:"email" validate:"email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserInput struct {
	Email    string `json:"email" validate:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
