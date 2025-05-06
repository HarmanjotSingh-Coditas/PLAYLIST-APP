package postgres

import (
	"fmt"
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
		viper.AddConfigPath("../../../src/config")
		viper.SetConfigName("postgres")
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("error retrieving database %w", err))
		}
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s password=%s user=%s sslmode=%s",
			viper.GetString("host"), viper.GetString("port"), viper.GetString("dbname"), viper.GetString("password"),
			viper.GetString("user"), viper.GetString("sslmode"))
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("error authenticating with the DB %w", err))
		}
	})
	if db == nil {
		panic("DB is nil")
	}
	return db
}
