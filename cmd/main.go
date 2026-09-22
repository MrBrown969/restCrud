package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrBrown969/restCrud/internal/adapter/controller/rest"
	"github.com/MrBrown969/restCrud/internal/adapter/repository"
	"github.com/MrBrown969/restCrud/internal/domain/usecase"
)

func main() {
	dbUrl := "postgres://crud_repo:12345678@localhost:5432/mobs?sslmode=disable"
	db, dbErr := repository.Connect(dbUrl)
	if dbErr != nil {
		log.Fatalf(fmt.Sprintf("error connecting to database: %s", dbErr))
	}

	repo := &repository.MobRepository{
		Db: db,
	}

	seeUsecase := usecase.SeeMobUsecase{Repo: repo}
	slayUsecase := usecase.SlayMobUsecase{Repo: repo}
	spawnUsecase := usecase.SpawnMobUsecase{Repo: repo}

	controller := rest.MobController{
		SeeUsecase:   &seeUsecase,
		SlayUsecase:  &slayUsecase,
		SpawnUsecase: &spawnUsecase,
	}

	server := http.NewServeMux()

	server.HandleFunc("GET /mob-pool/", controller.See)
	server.HandleFunc("DELETE /mob-pool/{id}", controller.Slay)
	server.HandleFunc("POST /mob-pool/", controller.Spawn)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println(err)
	}
}
