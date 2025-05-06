package models

type GenericAPIResponse struct {
	Message string
}

type BFFCreateUserPlaylistRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Song_ids    []int  `json:"song_ids" validate:"required,dive,gt=0"`
}
type BFFCreateUserPlaylistResponse struct {
	Message string `json:"message"`
}
