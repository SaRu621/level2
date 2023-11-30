package config

import (
	"flag"
	"strconv"
	"strings"
)

type Config struct {
	F []int
	D string
	S bool
}

func (cfg *Config) ParseArgs() {
	var buf string

	flag.BoolVar(&cfg.S, "s", false, "только строки с разделителем")
	flag.StringVar(&buf, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&cfg.D, "d", "\t", "использовать другой разделитель")

	flag.Parse()

	//args := flag.Args()

	splittedBuf := strings.Split(buf, ",")

	for i := range splittedBuf {
		value, err := strconv.Atoi(splittedBuf[i])

		if err != nil {
			cfg.F = []int{}
			break
			//os.Exit(2)
		}

		cfg.F = append(cfg.F, value)
	}
}
