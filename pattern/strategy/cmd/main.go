package main

import "strategy/pkg"

//поведенческий

var (
	start      = 10
	end        = 100
	strategies = []pkg.Strategy{
		&pkg.PublicTransportStrategy{},
		&pkg.RoadStrategy{},
		&pkg.WalkStrategy{},
	}
)

func main() {
	nav := pkg.Navigator{}

	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
