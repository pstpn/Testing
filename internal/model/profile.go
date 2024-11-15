package model

type Profile struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	City    string `json:"city"`
	Email   string `json:"email"`
}
