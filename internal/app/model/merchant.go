package model

type Merchant struct {
	ID string `json:"id" validate:"alphanum,max=256,required"`
	Type string `json:"type" validate:"required"`
	Name string `json:"name" validate:"alphanumdotunderscore,max=20,required"`
}