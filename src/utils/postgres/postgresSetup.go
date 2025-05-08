package postgres

import (
	"fmt"
	genericConstants "playlist-app/src/constants"
	genericModels "playlist-app/src/models"

	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetPostgresClient() *gorm.DB {
	once.Do(func() {
		viper.AddConfigPath(genericConstants.ConfigPath)
		viper.SetConfigName(genericConstants.ConfigName)
		viper.SetConfigType(genericConstants.ConfigType)
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf(genericConstants.DatabaseRetrievalError, err))
		}
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s password=%s user=%s sslmode=%s",
			viper.GetString(genericConstants.Host), viper.GetString(genericConstants.Port),
			viper.GetString(genericConstants.Dbname), viper.GetString(genericConstants.Password),
			viper.GetString(genericConstants.User), viper.GetString(genericConstants.Sslmode))
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf(genericConstants.DatabaseAuthenticationError, err))
		}
	})
	if db == nil {
		panic(genericConstants.DatabaseNilError)
	}
	db.AutoMigrate(&genericModels.Userss{}, &genericModels.Songs{}, &genericModels.Playlists{}, &genericModels.PlaylistSong{})
	return db
}
