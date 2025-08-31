package entity

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone_number"`
}
