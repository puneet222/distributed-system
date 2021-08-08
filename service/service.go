package service

import (
	"context"
	"distributed"
	"distributed/register"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func Start(ctx context.Context, service distributed.Service) (context.Context, error){
	ctx = startService(ctx, service)
	err := register.Add(service)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, service distributed.Service) context.Context {
	service.HttpHandleFunctions() // register http handles of service
	ctx, cancel := context.WithCancel(ctx)

	// create the server
	server := http.Server{}
	server.Addr = ":"+service.Port

	go func() {
		// start the server
		fmt.Printf("Started %v\n", service.Name)
		log.Println(server.ListenAndServe())
		fmt.Printf("Stopping %v", service.Name)
		cancel()
	}()

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		fmt.Println("Press control + C to quit")
		<- quit
		fmt.Printf("%v is shutting down", service.Name)
		register.Remove(service)
		server.Shutdown(ctx)
		cancel()
	}()
	return ctx
}