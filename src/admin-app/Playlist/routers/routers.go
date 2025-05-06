// package routers

// import (
// 	"admin-app/Playlist/business"
// 	"admin-app/Playlist/handler"
// 	"admin-app/Playlist/models"
// 	"admin-app/Playlist/repositiories"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"playlist-app/src/utils/postgres"

// 	"github.com/gin-gonic/gin"
// )

// func PookieLogger() gin.HandlerFunc {
// 	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
// 		// Assign cute status reactions 💅
// 		var vibe string
// 		switch {
// 		case param.StatusCode >= 500:
// 			vibe = "💥 Oopsie! Server's having a meltdown 😭"
// 		case param.StatusCode >= 400:
// 			vibe = "💔 Pookie messed up... try again bestie 🥺"
// 		case param.StatusCode >= 300:
// 			vibe = "🔁 Lil redirect detour hehe 💫"
// 		case param.StatusCode >= 200:
// 			vibe = "✨ All good bby, slayyy 💅"
// 		default:
// 			vibe = "🌈 Mysterious pookie moment 🤭"
// 		}

// 		// Return the final cutesy log string
// 		return fmt.Sprintf(
// 			"💖 [%s] ✨ From: %s 💌 Time: %v\n🍓 %s\n🔮 [%d %s] ➜ %s %s\n\n",
// 			param.TimeStamp.Format("2006/01/02 15:04:05"),
// 			param.ClientIP,
// 			param.Latency,
// 			vibe,
// 			param.StatusCode,
// 			http.StatusText(param.StatusCode),
// 			param.Method,
// 			param.Path,
// 		)
// 	})
// }

// func GetRouter() *gin.Engine {
// 	gin.SetMode(gin.ReleaseMode)
// 	router := gin.New()
// 	router.Use(PookieLogger(), gin.Recovery()) // 💅 Pookie logger in action

// 	db := postgres.GetPostgresClient()
// 	sqlDb, err := db.DB()
// 	if err != nil {
// 		panic(fmt.Errorf("error connecting to database %w", err))
// 	}
// 	err = sqlDb.Ping()
// 	if err != nil {
// 		panic(fmt.Errorf("error pinging database %w", err))
// 	}
// 	log.Println("🌟 Database? Connected, queen 👑")

// 	useDBMocks := false
// 	createUserPlaylistRepo := repositiories.GetCreateUserPlaylistRepository(useDBMocks)
// 	createUserPlaylistService := business.NewCreateUserPlaylistService(createUserPlaylistRepo)
// 	createUserPlaylistController := handler.NewCreateUserPlaylistController(*createUserPlaylistService)

// 	adSongsFromPlaylistRepo := repositiories.GetADSongsFromPlaylistRepository(useDBMocks)
// 	adSongsFromPlaylistService := business.NewAdSongsFromPlaylistService(adSongsFromPlaylistRepo)
// 	adSongsFromPlaylistController := handler.NewADSongsFromPlaylistController(*adSongsFromPlaylistService)

// 	v1 := router.Group("/v1")
// 	{
// 		v1.GET("/health", func(ctx *gin.Context) {
// 			ctx.JSON(http.StatusOK, models.GenericAPIResponse{
// 				Message: "💖 Pookie API is thriving ✨",
// 			})
// 		})
// 		v1.POST("/playlists", createUserPlaylistController.HandleCreateUserPlaylist)
// 		v1.PUT("/playlists", adSongsFromPlaylistController.HandleAdSongsFromPlaylist)
// 	}

// 	return router
// }

package routers

import (
	"admin-app/Playlist/business"
	"admin-app/Playlist/handler"
	"admin-app/Playlist/models"
	"admin-app/Playlist/repositiories"
	"fmt"
	"log"
	"net/http"
	"playlist-app/src/utils/postgres"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	db := postgres.GetPostgresClient()
	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("error connecting to database %w", err))
	}
	err = sqlDb.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging database %w", err))
	}
	log.Println("Connected to database successfully")

	useDBMocks := false
	createUserPlaylistRepo := repositiories.GetCreateUserPlaylistRepository(useDBMocks)
	createUserPlaylistService := business.NewCreateUserPlaylistService(createUserPlaylistRepo)
	createUserPlaylistController := handler.NewCreateUserPlaylistController(*createUserPlaylistService)

	adSongsFromPlaylistRepo := repositiories.GetADSongsFromPlaylistRepository(useDBMocks)
	adSongsFromPlaylistService := business.NewAdSongsFromPlaylistService(adSongsFromPlaylistRepo)
	adSongsFromPlaylistController := handler.NewADSongsFromPlaylistController(*adSongsFromPlaylistService)
	v1 := router.Group("/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, models.GenericAPIResponse{
				Message: "Service OK",
			})
		})
		v1.POST("/playlists", createUserPlaylistController.HandleCreateUserPlaylist)
		v1.PUT("/playlists", adSongsFromPlaylistController.HandleAdSongsFromPlaylist)

	}
	return router
}
