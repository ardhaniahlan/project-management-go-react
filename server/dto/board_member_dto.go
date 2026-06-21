package dto

type AddMembersRequest struct {
	UserIDs []string `json:"user_ids"`
}