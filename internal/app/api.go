// api/api.go

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// API is the base API server description
type API struct {
	// UNEXPORTED FIELDS
	config *Config
}

// New creates a new instance of the base API
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// Start initializes server loggers, router, database, etc.
// Start initializes server loggers, router, database, etc.
func (api *API) Start() error {
	fmt.Printf("Server is starting on port %s with logger level %s\n", api.config.Port, api.config.LoggerLevel)
	log.Println("Log message from Start function") // использование пакета log
	// Add your server initialization logic here


	// Setup your server components here
	router := http.NewServeMux()

	// Example route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	// Start the HTTP server
	err := http.ListenAndServe(api.config.Port, router)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	return nil
}

// ReadConfigFromFile reads the configuration from a JSON file
func ReadConfigFromFile(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}
