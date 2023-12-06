package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task8/pkg"
	"task8/pkg/config"
)

func main() {
	root := "/home/rus/"

	var cfg config.Config

	for {
		if root != "/home/rus/" {
			fmt.Print("user@user-ThinkPad-E14:~", root, "$ ")
		} else {
			fmt.Print("user@user-ThinkPad-E14:~$ ")
		}

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		input := scanner.Text()

		pipe := strings.Split(input, " | ")

		for _, command := range pipe {

			cfg.ParseConfig(command)

			switch cfg.Command {

			case "pwd":
				fmt.Println(root)

			case "ps":
				err := pkg.Ps()

				if err != nil {
					fmt.Println(err)
				}

			case "echo":
				for i := 1; i < len(cfg.Arg); i++ {
					fmt.Print(cfg.Arg[i], " ")
				}
				fmt.Println()

			case "cd":
				err := pkg.Cd(&root, cfg.Arg[1])

				if err != nil {
					fmt.Println(err)
					continue
				}

			case "kill":
				for i := 1; i < len(cfg.Arg); i++ {
					pid, err := strconv.Atoi(cfg.Arg[i])

					if err != nil {
						fmt.Println(err)
						continue
					}

					if pid == 0 {
						continue
					}

					err = pkg.Kill(pid)

					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}
