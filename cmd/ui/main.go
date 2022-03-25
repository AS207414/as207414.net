package main

import (
	"flag"
	"log"
	"os"
	"fmt"
)

var (
	version   string
)

type config struct {
	port int
	address string
}

type application struct {
	config config
    errorLog *log.Logger
    infoLog  *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "HTTP listen port")
	flag.StringVar(&cfg.address, "address", "0.0.0.0", "HTTP listening IP")

    displayVersion := flag.Bool("version", false, "Display version and exit")
	flag.Parse()

	if *displayVersion {
        fmt.Printf("version:\t%s\n", version)
        os.Exit(0)
    }

	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LUTC|log.Ltime|log.Llongfile)

	app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
    }

	infoLog.Printf("Starting server on http://%s:%d", cfg.address, cfg.port)
    err := app.serve()
    errorLog.Fatal(err)
}