package repositiories

import (
	"admin-app/Playlist/commons/constants"
	"context"
	"errors"
	"fmt"
	genericModels "playlist-app/src/models"

	"gorm.io/gorm"
)

type createUserPlaylistRepository struct{}

type CreateUserPlaylistRepository interface {
	CheckSongIdExists(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) (bool, error)
	CheckPlaylistExists(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) (bool, error)
	CreatePlaylist(ctx context.Context, db *gorm.DB, playlist genericModels.Playlist) (int, error)
	AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSong) error
}

func NewCreateUserPlaylistRepository() *createUserPlaylistRepository {
	return &createUserPlaylistRepository{}
}

func MockCreateUserPlaylistRepostiory() *createUserPlaylistRepository {
	return &createUserPlaylistRepository{}
}

func GetCreateUserPlaylistRepository(useDBMocks bool) CreateUserPlaylistRepository {
	if useDBMocks {
		return MockCreateUserPlaylistRepostiory()
	}
	return NewCreateUserPlaylistRepository()
}

func (repository *createUserPlaylistRepository) CheckSongIdExists(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) (bool, error) {
	var count int64
	err := db.WithContext(ctx).Model(&genericModels.Songs{}).Where(conditions).Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New(constants.SongIdDoesNotExistsError)
		}
		return false, fmt.Errorf(constants.SongIdExsistenceCheckingError, err)
	}
	return count > 0, nil
}

func (repository *createUserPlaylistRepository) CheckPlaylistExists(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) (bool, error) {
	var count int64
	err := db.WithContext(ctx).Model(&genericModels.Playlist{}).Where(conditions).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf(constants.PlaylisyExistenceCheckingError, err)
	}
	return count > 0, nil
}

func (repository *createUserPlaylistRepository) CreatePlaylist(ctx context.Context, db *gorm.DB, playlist genericModels.Playlist) (int, error) {
	if playlist.UserID == 0 {
		return 0, errors.New(constants.UserIdRequired)
	}
	err := db.WithContext(ctx).Create(&playlist).Error
	if err != nil {
		return 0, fmt.Errorf(constants.PlaylistCreationError, err)
	}
	return playlist.ID, nil
}

func (repository *createUserPlaylistRepository) AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSong) error {
	if len(playlistSongs) == 0 {
		return nil
	}
	err := db.WithContext(ctx).Create(&playlistSongs).Error
	if err != nil {
		return fmt.Errorf(constants.AddingSongsToPlaylistError, err)
	}
	return nil
}
