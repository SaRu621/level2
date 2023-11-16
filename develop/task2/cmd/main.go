package main

import (
	"fmt"
	"task2/pkg"
)

func main() {
	str := []string{"a4bc2d5e", "abcd", "45", ""}

	for _, val := range str {
		fmt.Println(string(pkg.Extraction(val)))
	}
}
