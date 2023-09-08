package structs

// date,precipitation,temp_max,temp_min,wind,weather
type WeatherLine struct {
	Date          string
	Precipitation float64
	TempMax       float64
	TempMin       float64
	Wind          float64
	Weather       string
}
