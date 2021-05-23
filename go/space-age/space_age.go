package space

type Planet string

func Age(input float64, planet Planet) float64 {
	var eys float64 = 31557600.0
	r := input / eys
	//var e float64
	var result float64

	switch planet {
	case "Earth":
		result = r / 1.0
	case "Mercury":
		result = r / 0.2408467
	case "Venus":
		result = r / 0.61519726
	case "Mars":
		result = r / 1.8808158
	case "Jupiter":
		result = r / 11.862615
	case "Saturn":
		result = r / 29.447498
	case "Uranus":
		result = r / 84.016846
	case "Neptune":
		result = r / 164.79132
	}
	return result
}
