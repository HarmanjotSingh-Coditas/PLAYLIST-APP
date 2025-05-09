package business

import (
	"admin-app/Playlist/commons/constants"
	"admin-app/Playlist/models"
	"admin-app/Playlist/repositiories"
	"context"
	"errors"
	genericModels "playlist-app/src/models"
	"playlist-app/src/utils/postgres"
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
	if len(BffAdSongsFromPlaylistRequest.Song_ids) == 0 {
		return nil, errors.New(constants.EmptySongIdsError)
	}

	db := postgres.GetPostgresClient()
	switch BffAdSongsFromPlaylistRequest.Action {
	case constants.Add:
		var songsToAdd []genericModels.PlaylistSong
		for _, songId := range BffAdSongsFromPlaylistRequest.Song_ids {
			songsToAdd = append(songsToAdd, genericModels.PlaylistSong{
				PlaylistID: BffAdSongsFromPlaylistRequest.PlaylistId,
				SongID:     songId,
			})
		}
		if err := service.repository.AddSongsToPlaylist(ctx, db, songsToAdd); err != nil {
			return nil, err
		}
	case constants.Delete:
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
