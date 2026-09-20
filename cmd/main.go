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

	controller := rest.MobController{SeeUsecase: &seeUsecase}

	server := http.NewServeMux()

	server.HandleFunc("/mob-pool/", controller.See)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println(err)
	}
}
