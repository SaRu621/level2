package main

import (
	"log"
	"task10/pkg"
	"task10/pkg/config"
)

func main() {
	var cfg config.Config

	if err := pkg.Telnet(&cfg); err != nil {
		log.Fatalln(err)
	}
}
