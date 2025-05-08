package business

import (
	"admin-app/Playlist/commons/constants"
	"admin-app/Playlist/models"
	"admin-app/Playlist/repositiories"
	"context"
	"errors"
	"fmt"
	genericModels "playlist-app/src/models"
	"playlist-app/src/utils/postgres"
	"strings"
)

type AdSongsFromPlaylistService struct {
	repository repositiories.ADSongsFromPlaylistRepository
}

func NewAdSongsFromPlaylistService(repository repositiories.ADSongsFromPlaylistRepository) *AdSongsFromPlaylistService {
	return &AdSongsFromPlaylistService{
		repository: repository,
	}
}

func (service *AdSongsFromPlaylistService) AdSongsPlaylistService(ctx context.Context, BffAdSongsFromPlaylistRequest models.BFFAdSongsFromPlaylistRequest) (*genericModels.Playlists, error) {
	db := postgres.GetPostgresClient()

	switch BffAdSongsFromPlaylistRequest.Action {
	case "ADD":
		var songsToAdd []genericModels.PlaylistSong
		for _, songId := range BffAdSongsFromPlaylistRequest.Song_ids {
			songsToAdd = append(songsToAdd, genericModels.PlaylistSong{
				PlaylistID: BffAdSongsFromPlaylistRequest.PlaylistId,
				SongID:     songId,
			})
		}

		// Let the database handle duplicate checks via unique constraint
		if err := service.repository.AddSongsToPlaylist(ctx, db, songsToAdd); err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				return nil, fmt.Errorf(constants.SongIdsAlreadyInPlaylistError)
			}
			return nil, err
		}

	case "DELETE":
		deleteConditions := map[string]interface{}{
			constants.PlaylistId: BffAdSongsFromPlaylistRequest.PlaylistId,
			constants.SongsId:    BffAdSongsFromPlaylistRequest.Song_ids,
		}
		if err := service.repository.DeleteSongsFromPlaylist(ctx, db, deleteConditions); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New(constants.InvalidAction)
	}

	updatedPlaylist, err := service.repository.GetPlaylistWithSongs(ctx, db, BffAdSongsFromPlaylistRequest.PlaylistId)
	if err != nil {
		return nil, err
	}
	return updatedPlaylist, nil
}
