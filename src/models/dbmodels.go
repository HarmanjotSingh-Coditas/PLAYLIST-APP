package models

import (
	"time"
)

type User struct {
	ID        int        `gorm:"primaryKey"`
	Playlists []Playlist `gorm:"foreignKey:UserID"`
}

type Songs struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"column:title"`
	Artist string `json:"artist" gorm:"column:artist"`
}

type Playlist struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	Songs       []Songs   `json:"songs" gorm:"many2many:playlist_songs;foreignKey:ID;joinForeignKey:playlist_id;References:ID;joinReferences:song_id"`
}

type PlaylistSong struct {
	PlaylistID int       `gorm:"primaryKey;column:playlist_id;not null"`
	SongID     int       `gorm:"primaryKey;column:song_id;not null"`
	AddedAt    time.Time `gorm:"column:added_at;autoCreateTime"`
}
