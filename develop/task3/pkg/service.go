package pkg

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"task3/pkg/structs"
)

func readFile(cfg structs.Config) ([][]string, error) { //Make [][]string from file
	file, err := os.Open(cfg.InputFile)

	if err != nil {
		fmt.Printf("File error: %s\n", err)
		return nil, err
	}

	defer file.Close()

	data := make([][]string, 0)

	scanner := bufio.NewScanner(file)

	unique := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()

		if cfg.Unique {
			if _, ok := unique[line]; !ok {
				unique[line] = true
				data = append(data, strings.Split(line, " "))
			} else {
				continue
			}
		} else {
			data = append(data, strings.Split(line, " "))
		}
	}

	return data, nil
}

func writeFile(cfg structs.Config, data [][]string) error {
	file, err := os.OpenFile(cfg.OutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	for _, val := range data {
		joinedLine := strings.Join(val, " ") + "\n"
		file.Write([]byte(joinedLine))
	}

	return nil
}

// ManSort -- сортировка файла
func ManSort(cfg structs.Config) error {
	//var cfg structs.Config

	//cfg.ParseArgs() // парсим аргументы из консоли

	data, err := readFile(cfg) // читаем файл

	if err != nil {
		return err
	}

	sort.Slice(data, func(i, j int) bool { // сортировка данных из файла
		if !cfg.NumSort {
			if !cfg.Reversed {
				return data[i][cfg.SortedColumn] < data[j][cfg.SortedColumn]
			}

			return data[i][cfg.SortedColumn] > data[j][cfg.SortedColumn]
		}
		if !cfg.Reversed {
			val1, err := strconv.Atoi(data[i][cfg.SortedColumn])

			if err != nil {
				return false
			}

			val2, err := strconv.Atoi(data[j][cfg.SortedColumn])

			if err != nil {
				return false
			}

			return val1 < val2
		}

		val1, err := strconv.Atoi(data[i][cfg.SortedColumn])

		if err != nil {
			return false
		}

		val2, err := strconv.Atoi(data[j][cfg.SortedColumn])

		if err != nil {
			return false
		}

		return val1 > val2

	})

	err = writeFile(cfg, data) // запись данных в файл

	if err != nil {
		return err
	}

	return nil

}
