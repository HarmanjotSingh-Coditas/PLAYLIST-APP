package routers

import (
	"admin-app/Playlist/business"
	"admin-app/Playlist/commons/constants"
	_ "admin-app/Playlist/docs" // Import generated Swagger docs
	"admin-app/Playlist/handler"
	"admin-app/Playlist/models"
	"admin-app/Playlist/repositiories"
	"fmt"
	"log"
	"net/http"
	"playlist-app/src/utils/postgres"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	db := postgres.GetPostgresClient()
	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf(constants.DatabaseConnectionError, err))
	}
	err = sqlDb.Ping()
	if err != nil {
		panic(fmt.Errorf(constants.DatabasePingingError, err))
	}
	log.Println(constants.DatabaseConnectionSuccess)

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	useDBMocks := false
	createUserPlaylistRepo := repositiories.GetCreateUserPlaylistRepository(useDBMocks)
	createUserPlaylistService := business.NewCreateUserPlaylistService(createUserPlaylistRepo)
	createUserPlaylistController := handler.NewCreateUserPlaylistController(*createUserPlaylistService)

	adSongsFromPlaylistRepo := repositiories.GetADSongsFromPlaylistRepository(useDBMocks)
	adSongsFromPlaylistService := business.NewAdSongsFromPlaylistService(adSongsFromPlaylistRepo)
	adSongsFromPlaylistController := handler.NewADSongsFromPlaylistController(*adSongsFromPlaylistService)
	v1 := router.Group(constants.Group)
	{
		v1.GET(constants.Health, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, models.GenericAPIResponse{
				Message: constants.ServiceOk,
			})
		})
		v1.POST(constants.CreatePlaylistRoute, createUserPlaylistController.HandleCreateUserPlaylist)
		v1.PUT(constants.AdPlaylistRoute, adSongsFromPlaylistController.HandleAdSongsFromPlaylist)

	}
	return router
}
