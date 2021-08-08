package main

import (
	"context"
	"distributed/logger"
	"distributed/service"
	"log"
)

func main() {
	loggerService := logger.NewLoggerService("app.log")
	ctx, err := service.Start(context.Background(), loggerService)
	if err != nil {
		log.Fatalln(err)
	}
	<- ctx.Done()
}
