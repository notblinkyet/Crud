package main

import (
	"fmt"
	"log"

	"net/http"
	"simple_crud/internal/config"
	"simple_crud/internal/services"
	"simple_crud/internal/storage"
	"simple_crud/internal/storage/posgresql"
)

var db storage.Storage

func main() {
	config, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	switch config.DataBase.Type {
	case "postgres":
		connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			config.DataBase.Username, config.DataBase.Password, config.DataBase.Host, config.DataBase.Port, config.DataBase.Name)
		db, err = posgresql.Open(connString)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("DataBase connect!")
		}

	default:
		log.Fatal("uknown db")
	}

	defer db.Close()

	server := services.New(config.Addr, http.NewServeMux(), db)
	server.FillEndpoits()
	log.Fatal(server.ListenAndServe())
}
