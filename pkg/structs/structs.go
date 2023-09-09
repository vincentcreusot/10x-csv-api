package structs

import "time"

// date,precipitation,temp_max,temp_min,wind,weather
type WeatherLine struct {
	Date          time.Time
	Precipitation float64
	TempMax       float64
	TempMin       float64
	Wind          float64
	Weather       string
	Corrupted     bool
}
