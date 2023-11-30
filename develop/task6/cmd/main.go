package main

import (
	"fmt"
	"task6/pkg"
	"task6/pkg/config"
)

func main() {
	var cfg config.Config

	cfg.ParseArgs()

	fmt.Println(cfg)

	pkg.ReadSTDIN(cfg)
}

/*

Имя;Возраст;Город
Анна;25;Москва
Иван;30;Санкт-Петербург
Мария;28;Киев
Николай;35
sample
*/
