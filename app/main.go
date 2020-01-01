package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/k0pernicus/go-tinyurl/internal/handlers"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (c Configuration) String() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func main() {
	// Read configuration file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Cannot read configuration file at root: %s", err.Error()))
	}
	var c Configuration
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal configuration error: %s", err.Error()))
	}

	// Register handlers
	fmt.Println("Registering handlers... ")
	router := mux.NewRouter()
	router.HandleFunc("/{id}", handlers.Get).Methods("GET")
	router.HandleFunc("/create", handlers.Create).Methods("POST")
	router.HandleFunc("/exists/{id}", handlers.Exists).Methods("GET")

	addr := c.String()

	// Create server
	srv := &http.Server{
		Handler: router,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server is running on %s\n", addr)

	// Serve
	srv.ListenAndServe()
}
