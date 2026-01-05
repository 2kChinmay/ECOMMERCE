package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2kChinmay/students-api/internal/config"
	student "github.com/2kChinmay/students-api/internal/http/handlers"
)

func main() {
	//lLOAD CONFIG:
	configuration := config.LoadConfigAndSerialize()

	//DB SETUP:

	//ROUTER SETUP:
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())

	//SERVER SETUP:
	server := &http.Server{
		Addr:    configuration.Http_server.Address,
		Handler: router,
	}

	slog.Info("Server started successfully", slog.String("address", `http://`+configuration.Http_server.Address))

	//Main function runs in a separate go routine and it can stop if any error occurs. To restrict it from stopping make a channel as follows:-
	channel := make(chan os.Signal, 1)
	//Notify into channel for the os terminate and interrupt signals
	signal.Notify(channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//This will run independently in a separate go routine
	go func() {
		//start server
		err := server.ListenAndServe()
		// if any err fuck the server
		if err != nil {
			log.Fatalf("Failed to start server %s", err.Error())
		}
	}()

	//teleport here if any signal (channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)  input occurs
	<-channel

	//GRACEFUL SHUTDOWN:
	slog.Info("Gracefully shutting down server")

	//context.Background() creates an empty point and container for storing context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//to cancel timeout after teleporting 	
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Error shutting down server:", slog.String("error",  err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
