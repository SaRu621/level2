package main

import (
	"task5/pkg"
	"task5/pkg/config"
)

func main() {
	var cfg config.Config

	cfg.ParseArgs()

	pkg.FindRows(cfg)
}
