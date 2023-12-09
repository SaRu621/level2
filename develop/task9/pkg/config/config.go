package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	Url *url.URL
	Rps int
	Rec bool
}

func (cfg *Config) ParseConfig() {
	var err error

	flag.BoolVar(&cfg.Rec, " r", false, "Рекурсивный проход по ссылкам <a>")
	flag.IntVar(&cfg.Rps, "rps", 1000, "кол-во запросов в секунду")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Не найден URL")
		os.Exit(2)
	}

	cfg.Url, err = url.Parse(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Не получилось распарсить URL: %s\n", err)
		os.Exit(2)
	}
}
