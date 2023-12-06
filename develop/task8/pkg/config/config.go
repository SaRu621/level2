package config

import (
	"os"
	"strings"
)

type Config struct {
	Command string
	Arg     []string
}

func (cfg *Config) ParseConfig(str string) {
	splittedStr := strings.Split(str, " ")

	cfg.Command = splittedStr[0]

	if cfg.Command == "pwd" || cfg.Command == "ps" {
		return
	}

	if len(splittedStr) < 2 {
		os.Exit(2)
	} else {
		cfg.Arg = []string{}
		cfg.Arg = append(cfg.Arg, splittedStr...)
	}
}
