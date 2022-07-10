package user

type CreateDto struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}

type UpdateDto struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}
