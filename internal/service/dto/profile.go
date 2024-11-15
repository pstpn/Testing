package dto

type CreateProfileRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	City    string `json:"city"`
	Email   string `json:"email"`
}

type GetProfileRequest struct {
	Email string `json:"email"`
}

type UpdateProfileRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	City    string `json:"city"`
	Email   string `json:"email"`
}

type DeleteProfileRequest struct {
	Email string `json:"email"`
}
