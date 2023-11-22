package structs

import (
	"flag"
	"fmt"
	"os"
)

// Config -- конфигурация команды (наличие ключей, файлов)
type Config struct {
	SortedColumn int  // номер колонки, по которой сортируем -k
	NumSort      bool // числовая сортировка -n
	Reversed     bool // Сортировка в обратном порядке -r
	Unique       bool // только уникальные строки -u

	InputFile  string
	OutputFile string
}

// ParseArgs -- метод, позволяющий собирать конфиг из данных консоли
func (cfg *Config) ParseArgs() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Инструкция %s", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&cfg.SortedColumn, "k", 0, "column for sort")
	flag.BoolVar(&cfg.NumSort, "n", false, "numeric sort")
	flag.BoolVar(&cfg.Reversed, "r", false, "desc sort")
	flag.BoolVar(&cfg.Unique, "u", false, "only unique rows")
	flag.Parse()

	flag.Usage()

	if cfg.SortedColumn < 0 {
		flag.Usage()
		os.Exit(2)
	}

	args := flag.Args()

	for _, val := range args {
		fmt.Println(val)
	}

	if len(args) > 2 { // Более двух аргументов -- выход из программы
		flag.Usage()
		os.Exit(2)
	}

	if len(args) >= 1 { // Более одного -- добавляем считываемый файл
		cfg.InputFile = args[0]
	}

	if len(args) == 2 { // Два аргумента -- запись в файл
		cfg.OutputFile = args[1]
	}

	check := func() func(bool) {
		var f bool
		return func(b bool) {
			if f && b {
				flag.Usage()
				os.Exit(2)
			}
			f = f || b
		}
	}()

	check(cfg.NumSort)
}
