package user

type ProfileResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateProfileRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}