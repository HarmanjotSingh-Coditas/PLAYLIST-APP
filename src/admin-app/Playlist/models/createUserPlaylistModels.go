package models

type GenericAPIResponse struct {
	Message string
}

type BFFCreateUserPlaylistRequest struct {
	UserID      uint16   `json:"userId" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Song_ids    []uint16 `json:"songIds" validate:"required"`
}
type BFFCreateUserPlaylistResponse struct {
	Message string `json:"message"`
}
