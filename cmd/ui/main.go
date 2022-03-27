package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

var (
	version string
	os_ver  string
	os_arc  string
	go_ver  string
	git_sha string
)

type config struct {
	port    int
	address string
}

type application struct {
	config        config
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "HTTP listen port")
	flag.StringVar(&cfg.address, "address", "", "HTTP listening IP")

	displayVersion := flag.Bool("version", false, "Display version and exit")
	flag.Parse()

	if *displayVersion {
		fmt.Printf("as207414-%s\n", version)
		fmt.Printf("%s/%s, %s, %s\n", os_ver, os_arc, go_ver, git_sha)
		os.Exit(0)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LUTC|log.Ltime|log.Llongfile)

	// Initialize a new template cache...
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		config:        cfg,
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}

	erra := app.serve()
	errorLog.Fatal(erra)
}
