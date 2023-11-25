package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task5/pkg/config"
)

func Check(a, b string, f bool) bool {
	if f {
		return a == b
	}

	return strings.Contains(a, b)
}

func FindRows(cfg config.Config) error {
	file, err := os.Open(cfg.File)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	i := 0

	counter := 0

	rowsFlags := make(map[int]bool, 0)
	rows := make([]string, 0)

	for scanner.Scan() {
		if _, ok := rowsFlags[i]; !ok {
			rowsFlags[i] = false
		}

		line := scanner.Text()

		rows = append(rows, line)

		if Check(line, cfg.ToFind, cfg.Fixed) || (Check(strings.ToLower(line), strings.ToLower(cfg.ToFind), cfg.Fixed) && cfg.IgnoreCase) {
			if cfg.Before > 0 {
				for j := 0; j <= cfg.Before; j++ {
					if i-j >= 0 {
						rowsFlags[i-j] = true
					}
				}
			}

			if cfg.After > 0 {
				for j := 0; j <= cfg.After; j++ {
					rowsFlags[i+j] = true
				}
			}

			if cfg.Context > 0 {
				for j := 0; j <= cfg.Context; j++ {
					rowsFlags[i+j] = true

					if i-j >= 0 {
						rowsFlags[i-j] = true
					}
				}
			}

			counter++

			rowsFlags[i] = true

		}
		i++
	}

	if cfg.Count {
		fmt.Println(counter)
	} else {
		for i, val := range rows {
			if (rowsFlags[i] && !cfg.Invert) || (!rowsFlags[i] && cfg.Invert) {
				fmt.Println(val)
			}
		}
	}

	return nil
}
