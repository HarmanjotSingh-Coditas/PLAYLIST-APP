package main

import (
	"admin-app/Playlist/routers"
	"log"
	"os"
	"os/signal"
)

func main() {
	router := routers.GetRouter()
	port := ":8080"

	log.Printf("Connecting server on port %s", port)
	go func() {
		if err := router.Run(port); err != nil {
			log.Fatalf("Unable to connect to server")
		}
	}()
	log.Printf("Connected server on port %s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Termination Signal Recieved , shutting down the server")

}
