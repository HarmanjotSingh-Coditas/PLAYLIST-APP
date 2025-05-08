package repositiories

import (
	"context"
	genericModels "playlist-app/src/models"

	"gorm.io/gorm"
)

type createUserPlaylistRepository struct{}

type CreateUserPlaylistRepository interface {
	CreatePlaylist(ctx context.Context, db *gorm.DB, playlist genericModels.Playlists) (uint16, error)
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

func (repository *createUserPlaylistRepository) CreatePlaylist(ctx context.Context, db *gorm.DB, playlist genericModels.Playlists) (uint16, error) {
	err := db.WithContext(ctx).Create(&playlist).Error
	if err != nil {
		return 0, err
	}
	if err == gorm.ErrDuplicatedKey {
		return 0, gorm.ErrDuplicatedKey
	}
	if err == gorm.ErrForeignKeyViolated {
		return 0, gorm.ErrForeignKeyViolated
	}
	return playlist.ID, nil
}

func (repository *createUserPlaylistRepository) AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSong) error {
	if len(playlistSongs) == 0 {
		return nil
	}
	err := db.WithContext(ctx).Create(&playlistSongs).Error

	if err != nil {
		return err
	}
	return nil
}
