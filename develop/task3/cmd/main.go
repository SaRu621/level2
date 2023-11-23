package main

import (
	"task3/pkg"
	"task3/pkg/structs"
)

// /home/rus/GoFolder/level2/develop/task3/input.txt

func main() {
	var cfg structs.Config

	cfg.ParseArgs()

	pkg.ManSort(cfg)
}
