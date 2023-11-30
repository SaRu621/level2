package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task6/pkg/config"
)

func ReadSTDIN(cfg config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	responseOut := ""

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			break
		}

		if !strings.Contains(input, cfg.D) && cfg.S {
			continue
		}

		parsedInput := strings.Split(input, cfg.D)

		for _, val := range cfg.F {
			if val-1 >= len(parsedInput) {
				break
			}

			responseOut += parsedInput[val-1] + " "
		}

		responseOut += "\n"
	}

	fmt.Println(responseOut)

	os.Exit(0)
}
