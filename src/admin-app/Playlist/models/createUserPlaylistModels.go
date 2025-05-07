package models

type GenericAPIResponse struct {
	Message string
}

type BFFCreateUserPlaylistRequest struct {
	UserID      int    `json:"userId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Song_ids    []int  `json:"songIds" validate:"required,dive,gt=0"`
}
type BFFCreateUserPlaylistResponse struct {
	Message string `json:"message"`
}
