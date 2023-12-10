package config

import (
	"flag"
	"log"
	"strconv"
	"time"
)

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func (cfg *Config) ParseConfig() {
	flag.DurationVar(&cfg.Timeout, "timeout", 10*time.Second, "устанавливает время таймаута.")
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		log.Fatalln("len(args) != 2")
	}

	cfg.Host = args[0]

	var err error

	cfg.Port, err = strconv.Atoi(args[1])
	if err != nil {
		flag.Usage()
		log.Fatalln("Порт -- это число")
	}
	if cfg.Port < 1 || cfg.Port > 65535 {
		log.Fatalln("Порт должн быть в диапазоне [1, 65535]")
	}
}
