package main

import (
	"fmt"
	"log"
	"net/http"

	"clean_architecture/pkg/store/postgres"
	delivery "clean_architecture/services/contact/internal/delivery/http"
	repository "clean_architecture/services/contact/internal/repository/storage/postgres"
	contactUseCase "clean_architecture/services/contact/internal/useCase/contact"
	groupUseCase "clean_architecture/services/contact/internal/useCase/group"
)

func main() {
	config := postgres.NewDBConfig("localhost", 5432, "postgres", "123321", "postgres")
	db, err := postgres.ConnectToDB(config)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	r := repository.New(db)
	ucGroup := groupUseCase.New(r)
	ucContact := contactUseCase.New(r)

	d := delivery.New(ucContact, ucGroup)

	addr := 4000
	addrStr := fmt.Sprintf(":%d", addr)

	log.Printf("Starting server on port: %d", addr)

	if err := http.ListenAndServe(addrStr, d.Router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
