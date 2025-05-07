package models

import genericModels "playlist-app/src/models"

type BFFAdSongsFromPlaylistRequest struct {
	Action     string `json:"action" validate:"required" example:"ADD" enums:"ADD,DELETE"`
	PlaylistId int    `json:"playlistId" validate:"required" example:"1"`
	Song_ids   []int  `json:"songIds" validate:"required" example:"1,2,3"`
	UserID     int    `json:"userId" validate:"required" example:"1"`
}

type BFFAdSongsFromPlaylistResponse struct {
	Message  string                 `json:"message" example:"Songs added to playlist successfully"`
	Playlist genericModels.Playlist `json:"playlist" swaggertype:"object"`
}
