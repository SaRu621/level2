package main

import (
	"builder/pkg"
)

func main() {
	asusCollector := pkg.GetCollector("asus")
	hpCollector := pkg.GetCollector("hp")

	builder := pkg.NewBuilder(asusCollector)

	asusComp := builder.CreateComputer()

	asusComp.Print()

	builder.SetCollector(hpCollector)

	hpComp := builder.CreateComputer()

	hpComp.Print()
}
