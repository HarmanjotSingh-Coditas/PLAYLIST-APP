package business

import (
	"admin-app/Playlist/models"
	"admin-app/Playlist/repositiories"
	"context"
	"errors"
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

	for _, songID := range bffCreateUserPlaylist.Song_ids {
		conditions := map[string]interface{}{
			"id": songID,
		}
		columns := []string{"id"}
		exists, err := service.repository.CheckSongIdExists(ctx, db, columns, conditions)
		if err != nil {
			return false, err
		}
		if !exists {
			return false, errors.New("one or more song IDs do not exist")
		}
	}

	conditions := map[string]interface{}{
		"name": bffCreateUserPlaylist.Name,
	}
	columns := []string{"name"}
	exists, err := service.repository.CheckPlaylistExists(ctx, db, columns, conditions)
	if err != nil {
		return false, err
	}
	if exists {
		return false, errors.New("playlist already exists")
	}

	playlist := genericModels.Playlist{
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
			PlaylistID: playlistID,
			SongID:     songID,
		})
	}

	err = service.repository.AddSongsToPlaylist(ctx, db, playlistSongs)
	if err != nil {
		return false, err
	}

	return true, nil
}
