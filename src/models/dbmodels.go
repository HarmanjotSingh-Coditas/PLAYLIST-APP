package models

import "time"

// type User struct {
// 	ID        int        `gorm:"primaryKey"`
// 	Playlists []Playlist `gorm:"foreignKey:UserID"`
// }

// // type Songs struct {
// // 	ID     int    `json:"id" gorm:"primaryKey"`
// // 	Title  string `json:"title" gorm:"column:title"`
// // 	Artist string `json:"artist" gorm:"column:artist"`
// // }

// // type Playlist struct {
// // 	ID          int       `json:"id" gorm:"primaryKey"`
// // 	UserID      int       `json:"user_id" gorm:"column:user_id"`
// // 	Name        string    `json:"name" gorm:"column:name"`
// // 	Description string    `json:"description" gorm:"column:description"`
// // 	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
// // 	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
// // 	Songs       []Songs   `json:"songs" gorm:"many2many:playlist_songs;foreignKey:ID;joinForeignKey:playlist_id;References:ID;joinReferences:song_id"`
// // }

// // type PlaylistSong struct {
// // 	PlaylistID int       `gorm:"primaryKey;column:playlist_id;not null"`
// // 	SongID     int       `gorm:"primaryKey;column:song_id;not null"`
// // 	AddedAt    time.Time `gorm:"column:added_at;autoCreateTime"`
// // }

type Userss struct {
	ID uint16 `gorm:"primaryKey;column:id;autoIncrement"`
}

type Songs struct {
	ID     uint16 `gorm:"primaryKey;column:id;autoIncrement"`
	Title  string `gorm:"column:title;not null"`
	Artist string `gorm:"column:artist;not null"`
}

type Playlists struct {
	ID          uint16    `gorm:"primaryKey;column:id;autoIncrement"`
	UserID      uint16    `gorm:"column:user_id;uniqueIndex:idx_user_playlist;not null"`
	Name        string    `gorm:"column:name;uniqueIndex:idx_user_playlist;not null"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	Songs       []Songs   `gorm:"many2many:playlist_songs;foreignKey:ID;joinForeignKey:playlist_id;References:ID;joinReferences:song_id"`
}

type PlaylistSong struct {
	PlaylistID uint16    `gorm:"column:playlist_id;not null;uniqueIndex:idx_playlist_song"`
	SongID     uint16    `gorm:"column:song_id;not null;uniqueIndex:idx_playlist_song"`
	AddedAt    time.Time `gorm:"column:added_at;default:CURRENT_TIMESTAMP"`
	Playlist   Playlists `gorm:"foreignKey:PlaylistID;refrences:ID;constraint:OnUpdate:CASCADE"`
	Song       Songs     `gorm:"foreignKey:SongID;refrences:ID;constraint:OnUpdate:CASCADE"`
}
