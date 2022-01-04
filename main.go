package main

import (
	"log"
	"flag"
)

const version = "1.0.0"

// Config struct
type Config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type Application struct {
	config Config
	logger *log.Logger
}

func main() {

	// setting env variables manually (improvement: will use go dotenv external package)
	// os.Setenv("DSN", "KASIMUDEY")

	// getting env variable
	// fmt.Println(os.Getenv("DSN"))

	// declare an instance of the config struct
	var cfg Config

	// initializing config varibles
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

}
