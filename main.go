package Navigator

import "Navigator/pkg"

var (
	start      = 10
	end        = 112
	strategies = []pkg.Strategy{
		&pkg.Bus{},
		&pkg.RoadStrategy{},
		&pkg.Walk{},
	}
)

func main() {
	nav := pkg.Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)

	}
}
