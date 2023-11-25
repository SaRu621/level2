package config

import (
	"flag"
	"os"
)

type Config struct {
	After   int //-A flag
	Before  int //-B flag
	Context int //-C = (-A + -B) flag

	Count      bool //-c flag
	IgnoreCase bool //-i flag
	Invert     bool //-v flag
	Fixed      bool //-F flag
	Numeric    bool //-n flag

	ToFind string
	File   string
}

func (cfg *Config) ParseArgs() {
	flag.IntVar(&cfg.After, "A", 0, "Количество строк перед найденной строкой")
	flag.IntVar(&cfg.Before, "B", 0, "Количество строк после найденной строкой")
	flag.IntVar(&cfg.Context, "C", 0, "Количество строк после и перед найденной строкой")
	
	flag.BoolVar(&cfg.Count, "c", false, "Количество найденных строк")
	flag.BoolVar(&cfg.IgnoreCase, "i", false, "Игнорирование регистра")
	flag.BoolVar(&cfg.Invert, "v", false, "Вывод строк, не удовлетворяющих условию")
	flag.BoolVar(&cfg.Fixed, "F", false, "Количество строк, в которых найдено совпадение")

	flag.Parse()

	args := flag.Args()

	cfg.ToFind = args[0]
	cfg.File = args[1]

	if len(args) != 2 {
		os.Exit(2)
	}
}
