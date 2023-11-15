package pkg

type Triangle struct {
	Side int
}

func (t *Triangle) Accept(v Visitor) {
	v.VisitForTriangle(t)
}

func (t *Triangle) GetType() string {
	return "Triangle"
}
