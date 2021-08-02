package main

import (
	"log"
	"rtforum/config"
	"rtforum/internal/api/handler"
	"rtforum/internal/api/server"
	"rtforum/internal/repository"
	"rtforum/internal/usecase"
)

func main() {
	srv := new(server.Server)
	if err := srv.Run(config.PORT, handler.Router()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()
	repo := repository.NewRepository(db)
	deps := &usecase.UseCaseDeps{
		Repo: repo,
	}
	useCases := usecase.NewUseCases(deps)
}
