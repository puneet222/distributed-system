package main

import (
	"context"
	"distributed/register"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	registerService := register.NewRegistrationService()
	registerService.HttpHandleFunctions() // register http handles of service

	// create the server
	server := http.Server{}
	server.Addr = ":"+registerService.Port

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		// start the server
		fmt.Printf("Started %v\n", registerService.Name)
		log.Println(server.ListenAndServe())
		fmt.Printf("Stopping %v", registerService.Name)
		cancel()
	}()

	go func() {
		quit := make(chan os.Signal)
		fmt.Println("Press control + C to quit")
		<- quit
		server.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	log.Println("Shutting down registry service")
}
