package dto

type NewListRequest struct {
	Title         string `json:"title" validate:"required"`
	BoardPublicID string `json:"board_public_id" validate:"required"`
}

type UpdateListRequest struct {
	Title    string `json:"title" validate:"required"`
}
