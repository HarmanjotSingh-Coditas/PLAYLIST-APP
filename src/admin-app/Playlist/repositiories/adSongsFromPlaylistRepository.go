package repositiories

import (
	"admin-app/Playlist/commons/constants"
	"context"

	genericModels "playlist-app/src/models"

	"gorm.io/gorm"
)

type adSongsFromPlaylistRepository struct{}

type ADSongsFromPlaylistRepository interface {
	AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSong) error
	DeleteSongsFromPlaylist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error
	GetPlaylistWithSongs(ctx context.Context, db *gorm.DB, playlistID uint16) (*genericModels.Playlists, error)
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

func (repository *adSongsFromPlaylistRepository) AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSong) error {
	err := db.WithContext(ctx).Create(&playlistSongs).Error
	return err
}

func (repository *adSongsFromPlaylistRepository) DeleteSongsFromPlaylist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error {
	result := db.WithContext(ctx).Where(conditions).Delete(&genericModels.PlaylistSong{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repository *adSongsFromPlaylistRepository) GetPlaylistWithSongs(ctx context.Context, db *gorm.DB, playlistID uint16) (*genericModels.Playlists, error) {
	var playlist genericModels.Playlists
	err := db.WithContext(ctx).
		Preload(constants.Songs).
		Where(constants.WehereIdClause, playlistID).
		First(&playlist).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}
