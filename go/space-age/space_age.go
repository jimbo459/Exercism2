package space

func Age(seconds int64, planet string) float64 {

	/* Map containing all the planet  */
	var planetInfo = map[string]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	/* Define earth Year seconds */
	const earthYear = 31557600

	/* Calculate planetYear in seconds */
	planetYear := earthYear * planetInfo[planet]

	/* Return years old */
	return float64(seconds) / planetYear
}
