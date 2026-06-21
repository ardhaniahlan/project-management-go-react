package dto

type AddMembersRequest struct {
	UserIDs []string `json:"user_ids"`
}

type RemoveMembersRequest struct {
	UserIDs []string `json:"user_ids"`
}