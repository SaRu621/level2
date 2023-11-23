package main

import (
	"fmt"
	"task4/pkg"
)

func main() {
	words := []string{"Пятка", "Тяпка", "пятак", "листок", "столик", "слиток"}

	t := pkg.Anagram(&words)

	fmt.Println(t)
}
