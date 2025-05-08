package business

import (
	"admin-app/Playlist/models"
	"admin-app/Playlist/repositiories"
	"context"
	genericModels "playlist-app/src/models"
	"playlist-app/src/utils/postgres"
)

type CreateUserPlaylistService struct {
	repository repositiories.CreateUserPlaylistRepository
}

func NewCreateUserPlaylistService(repository repositiories.CreateUserPlaylistRepository) *CreateUserPlaylistService {
	return &CreateUserPlaylistService{
		repository: repository,
	}
}

func (service *CreateUserPlaylistService) CreateUserPlaylistService(ctx context.Context, bffCreateUserPlaylist models.BFFCreateUserPlaylistRequest) (bool, error) {
	db := postgres.GetPostgresClient()

	playlist := genericModels.Playlists{
		UserID:      bffCreateUserPlaylist.UserID,
		Name:        bffCreateUserPlaylist.Name,
		Description: bffCreateUserPlaylist.Description,
	}

	playlistID, err := service.repository.CreatePlaylist(ctx, db, playlist)
	if err != nil {
		return false, err
	}

	var playlistSongs []genericModels.PlaylistSong
	for _, songID := range bffCreateUserPlaylist.Song_ids {
		playlistSongs = append(playlistSongs, genericModels.PlaylistSong{
			PlaylistID: uint16(playlistID),
			SongID:     songID,
		})
	}

	err = service.repository.AddSongsToPlaylist(ctx, db, playlistSongs)
	if err != nil {
		return false, err
	}

	return true, nil
}
