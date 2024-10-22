package main

import (
	"fmt"
	"log"

	"net/http"
	"simple_crud/pkg/config"
	"simple_crud/pkg/services"
	"simple_crud/pkg/storage"
	"simple_crud/pkg/storage/posgresql"
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
