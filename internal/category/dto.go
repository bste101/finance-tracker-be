package category

type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Type  string `json:"type" binding:"required,oneof=income expense"`
	Color string `json:"color"`
	Icon  string `json:"icon"`
}

type UpdateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Type  string `json:"type" binding:"required,oneof=income expense"`
	Color string `json:"color"`
	Icon  string `json:"icon"`
}

type CategoryResponse struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"userId"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Color  string `json:"color"`
	Icon   string `json:"icon"`
}
