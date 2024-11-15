package http

type StatusBadRequest struct {
	Error string `json:"error" example:"Incorrect request body"`
}

type StatusUnauthorized struct {
	Error string `json:"error" example:"Cant login user"`
}

type StatusNotFound struct {
	Error string `json:"error" example:"Failed to get profile"`
}

type StatusForbidden struct {
	Error string `json:"error" example:"Incorrect user role"`
}
