package pkg

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForTriangle(*Triangle)
}
