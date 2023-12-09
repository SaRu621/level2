package main

import (
	"log"
	"os"
	"task9/pkg"
	"task9/pkg/config"
)

func main() {
	var cfg config.Config
	cfg.ParseConfig()

	logger := log.New(os.Stderr, "error: ", log.Ldate|log.Ltime|log.Lshortfile)

	pkg.Wget(&cfg, logger)
}
