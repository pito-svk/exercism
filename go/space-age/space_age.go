package space

// Planet name
type Planet string

// Age calculates how old someone would be on specific planet
func Age(seconds float64, planet Planet) float64 {
	var planetOffset float64

	switch planet {
	case "Earth":
		planetOffset = 1
	case "Mercury":
		planetOffset = 0.2408467
	case "Venus":
		planetOffset = 0.61519726
	case "Mars":
		planetOffset = 1.8808158
	case "Jupiter":
		planetOffset = 11.862615
	case "Saturn":
		planetOffset = 29.447498
	case "Uranus":
		planetOffset = 84.016846
	case "Neptune":
		planetOffset = 164.79132
	}

	return seconds / 60 / 60 / 24 / 365.25 / planetOffset
}
