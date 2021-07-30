package model

type User struct {
	ID        string `json:"id" validate:"alphanum,max=256,required"`
	Type      string `json:"type" validate:"required"`
	UserName  string `json:"userName" validate:"alphanumdotunderscore,max=20,required"`
	FirstName string `json:"firstName" validate:"max=100"`
	LastName  string `json:"lastName" validate:"max=100"`
	Email     string `json:"email" validate:"omitempty,email"`
}
