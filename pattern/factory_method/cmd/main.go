package main

import "factory_method/pkg"

var types = []string{pkg.PersonalComputerType, pkg.NotebookType, pkg.ServerType, "monoblock"}

func main() {
	for _, val := range types {
		computer := pkg.New(val)

		if computer != nil {
			computer.PrintDetails()
		}
	}
}
