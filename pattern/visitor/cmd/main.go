package main

import (
	"fmt"
	"visitor/pkg"
)

type AreaCalc struct {
	area float64
}

func (*AreaCalc) VisitForSquare(*pkg.Square) {
	fmt.Println("Calc area for square")
}

func (*AreaCalc) VisitForTriangle(*pkg.Triangle) {
	fmt.Println("Calc area for triangle")
}

func (*AreaCalc) VisitForCircle(*pkg.Circle) {
	fmt.Println("Calc area for cirlce")
}

func main() {
	square := &pkg.Square{Side: 2}

	circle := &pkg.Circle{Radius: 3}

	triangle := &pkg.Triangle{Side: 1}

	areaCalc := &AreaCalc{}

	square.Accept(areaCalc)

	circle.Accept(areaCalc)

	triangle.Accept(areaCalc)

}
