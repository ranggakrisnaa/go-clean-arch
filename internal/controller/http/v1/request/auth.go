package request

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Register struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
	FullName string `json:"fullname" validate:"required"`
	Phone    string `json:"phone_number"    validate:"required"`
}
