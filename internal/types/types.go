package types

type User struct {
	ID string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required"`
	Age int `json:"age" validate:"required,min=18,max=100"`
}