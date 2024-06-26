package models

type User struct {
	ID int `json:"id" example:"1"`
	Name string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password,omitempty"`
}