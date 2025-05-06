package business

import (
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

func (service *AdSongsFromPlaylistService) AdSongsPlaylistService(ctx context.Context, BffAdSongsFromPlaylistRequest models.BFFAdSongsFromPlaylistRequest) (*genericModels.Playlist, error) {
	db := postgres.GetPostgresClient()

	checkPlaylistExistCondition := map[string]interface{}{
		"id": BffAdSongsFromPlaylistRequest.PlaylistId,
	}

	exists, err := service.repository.CheckPlaylistExists(ctx, db, checkPlaylistExistCondition)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("playlist does not exist")
	}

	switch BffAdSongsFromPlaylistRequest.Action {
	case "ADD":
		var songsToAdd []genericModels.PlaylistSong
		var existingSongs []int

		for _, songId := range BffAdSongsFromPlaylistRequest.Song_ids {
			dupExistenceCheck := map[string]interface{}{
				"playlist_id": BffAdSongsFromPlaylistRequest.PlaylistId,
				"song_id":     songId,
			}
			exists, err = service.repository.CheckSongsExistsInPlaylist(ctx, db, dupExistenceCheck)
			if err != nil {
				continue
			}
			if !exists {
				songsToAdd = append(songsToAdd, genericModels.PlaylistSong{
					PlaylistID: BffAdSongsFromPlaylistRequest.PlaylistId,
					SongID:     songId,
				})
			} else {
				existingSongs = append(existingSongs, songId)
			}
		}

		if len(songsToAdd) == 0 {
			return nil, fmt.Errorf("songs with IDs [%v] already exist in playlist", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(existingSongs)), ", "), "[]"))
		}

		err = service.repository.AddSongsToPlaylist(ctx, db, songsToAdd)
		if err != nil {
			return nil, err
		}

	case "DELETE":
		var songsToDelete []genericModels.PlaylistSong
		var nonExistentSongs []int

		for _, songID := range BffAdSongsFromPlaylistRequest.Song_ids {
			existenceCheck := map[string]interface{}{
				"playlist_id": BffAdSongsFromPlaylistRequest.PlaylistId,
				"song_id":     songID,
			}
			exists, err = service.repository.CheckSongsExistsInPlaylist(ctx, db, existenceCheck)
			if err != nil {
				continue
			}
			if exists {
				songsToDelete = append(songsToDelete, genericModels.PlaylistSong{
					PlaylistID: BffAdSongsFromPlaylistRequest.PlaylistId,
					SongID:     songID,
				})
			} else {
				nonExistentSongs = append(nonExistentSongs, songID)
			}
		}

		if len(songsToDelete) == 0 {
			return nil, fmt.Errorf("no valid songs to delete. Non-existent songs: %v", nonExistentSongs)
		}

		for _, songToDelete := range songsToDelete {
			deleteConditions := map[string]interface{}{
				"playlist_id": songToDelete.PlaylistID,
				"song_id":     songToDelete.SongID,
			}
			if err := service.repository.DeleteSongsFromPlaylist(ctx, db, deleteConditions); err != nil {
				continue
			}
		}

	default:
		return nil, errors.New("invalid action")
	}

	updatedPlaylist, err := service.repository.GetPlaylistWithSongs(ctx, db, BffAdSongsFromPlaylistRequest.PlaylistId)
	if err != nil {
		return nil, err
	}
	return updatedPlaylist, nil
}
