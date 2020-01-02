package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	app "github.com/k0pernicus/go-tinyurl/internal"
	"github.com/k0pernicus/go-tinyurl/internal/handlers"
	"gopkg.in/yaml.v2"
)

func main() {
	// Read configuration file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Cannot read configuration file at root: %s", err.Error()))
	}
	err = yaml.Unmarshal(yamlFile, &app.C)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal configuration error: %s", err.Error()))
	}

	// Connect to sqlite DB
	if err = app.ConnectDB("../db/urls.db"); err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %s\n", err.Error()))
	}
	defer app.DB.Close()

	// Register handlers
	fmt.Println("Registering handlers... ")
	router := mux.NewRouter()
	router.HandleFunc("/{id}", handlers.Get).Methods("GET")
	router.HandleFunc("/create", handlers.Create).Methods("POST")
	router.HandleFunc("/exists/{id}", handlers.Exists).Methods("GET")

	addr := app.C.String()

	// Create server
	srv := &http.Server{
		Handler: router,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	defer srv.Close()

	fmt.Printf("Server is running on %s\n", addr)

	// Serve
	srv.ListenAndServe()
}
