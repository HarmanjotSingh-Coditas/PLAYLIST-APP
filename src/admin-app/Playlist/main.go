package main

import (
	"admin-app/Playlist/commons/constants"
	"admin-app/Playlist/routers"

	"log"
	"os"
	"os/signal"
	genericConstants "playlist-app/src/constants"
)

func main() {
	router := routers.GetRouter()
	port := constants.Port

	log.Printf(genericConstants.ConnectingToServer, port)
	go func() {
		if err := router.Run(port); err != nil {
			log.Fatalf(genericConstants.UnableToConnectServerError)
		}
	}()
	log.Printf(genericConstants.ConnectedToServer, port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println(genericConstants.ServerShutdownSignal)

}
