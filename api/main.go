package main

import (
	"flag"
	"fmt"
	"log"
	"context"
	"database/sql"
	"os"
	"net/http"
	"time"
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
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("DSN"), "PostgreSQL DSN")

	flag.Parse()

	// Initialize a new logger which writes messages to the standard out stream
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// connecto to db
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// defer a call to close the connection
	defer db.Close()

	// log msg upon successful db connection
	logger.Printf("database connection pool established")

	// an instance of the Application struct
	app := Application{
		config: cfg,
		logger: logger,
	}

	// use the httprouter instance returned by app.routes() as the server handler.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	// Start the HTTP
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}



/**
* function to open db connection
* reurns sql.DB connection pool, error
*/
func openDB(cfg Config) (*sql.DB, error){
	// create empty connection pool
	db, err := sql.Open("postgres", cfg.db.dsn)

	if err != nil{
		return nil, err
	}

	// create a context with a 5 sec timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// use PingContext to establish a new connection to the database
	err = db.PingContext(ctx)
	if err != nil{
		return nil, err
	}

	// return the sql.DB connection pool
	return db, nil
}

