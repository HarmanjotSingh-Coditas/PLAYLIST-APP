package repositiories

import (
	"admin-app/Playlist/commons/constants"
	"context"
	"errors"

	genericModels "playlist-app/src/models"

	"gorm.io/gorm"
)

type adSongsFromPlaylistRepository struct{}

type ADSongsFromPlaylistRepository interface {
	CheckPlaylistExists(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (bool, error)
	CheckSongsExistsInPlaylist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (bool, error)
	AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSong) error
	DeleteSongsFromPlaylist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error
	GetPlaylistWithSongs(ctx context.Context, db *gorm.DB, playlistID int) (*genericModels.Playlist, error)
}

func NewADSongsFromPlaylistRepository() *adSongsFromPlaylistRepository {
	return &adSongsFromPlaylistRepository{}
}

func MockADSongsFromPlaylistRepository() *adSongsFromPlaylistRepository {
	return &adSongsFromPlaylistRepository{}
}

func GetADSongsFromPlaylistRepository(useDBmocks bool) ADSongsFromPlaylistRepository {
	if useDBmocks {
		return MockADSongsFromPlaylistRepository()
	}
	return NewADSongsFromPlaylistRepository()
}

func (repository *adSongsFromPlaylistRepository) CheckPlaylistExists(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (bool, error) {
	var count int64
	err := db.WithContext(ctx).Model(&genericModels.Playlist{}).Where(conditions).Count(&count).Error
	if err != nil {
		return false, errors.New(constants.PlaylistExistenceCheckError)
	}
	return count > 0, nil
}

func (repository *adSongsFromPlaylistRepository) CheckSongsExistsInPlaylist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (bool, error) {
	var count int64
	err := db.WithContext(ctx).Model(&genericModels.PlaylistSong{}).Where(conditions).Count(&count).Error
	if err != nil {
		return false, errors.New(constants.SongExistenceCheckingError)
	}
	return count > 0, nil
}

func (repository *adSongsFromPlaylistRepository) AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSong) error {
	if len(playlistSongs) == 0 {
		return errors.New(constants.NoSongsProvidedError)
	}
	err := db.WithContext(ctx).Create(&playlistSongs).Error
	if err != nil {
		return errors.New(constants.SongAddingToPlaylistError)
	}
	return nil
}

func (repository *adSongsFromPlaylistRepository) DeleteSongsFromPlaylist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error {
	result := db.WithContext(ctx).Model(&genericModels.PlaylistSong{}).Where(conditions).Delete(&genericModels.PlaylistSong{})
	if result.Error != nil {
		return errors.New(constants.SongDeletionfromPlaylistError)
	}
	if result.RowsAffected == 0 {
		return errors.New(constants.NoSongsDeletedFromPlaylist)
	}
	return nil
}

func (repository *adSongsFromPlaylistRepository) GetPlaylistWithSongs(ctx context.Context, db *gorm.DB, playlistID int) (*genericModels.Playlist, error) {
	var playlist genericModels.Playlist
	err := db.WithContext(ctx).
		Preload(constants.Songs).
		Where(constants.WehereIdClause, playlistID).
		First(&playlist).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}
