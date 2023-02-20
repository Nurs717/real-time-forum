package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"rtforum/config"
	"rtforum/internal/api/rest"
	"rtforum/internal/api/server"
	"rtforum/internal/repository"
	"rtforum/internal/usecase"
	"syscall"
	"time"
)

func main() {
	srv := new(server.Server)
	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()
	repo := repository.NewRepository(db, time.Second*3)
	deps := &usecase.UseCaseDeps{
		Repo: repo,
	}
	useCases := usecase.NewUseCases(deps)
	handlers := rest.NewHandler(useCases)
	go func() {
		if err := srv.Run(config.API_PORT, handlers.Router()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Printf("Server has started at port: %s\n", config.API_PORT)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	log.Printf("server is stopping at port: %s\n", config.API_PORT)
	if err := srv.ShutDown(context.Background()); err != nil {
		log.Fatalf("error occured while shutting down server: %s\n", err.Error())
	}
}
