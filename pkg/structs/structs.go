package structs

// date,precipitation,temp_max,temp_min,wind,weather
type WeatherLine struct {
	Date          string  `json:"date"`
	Precipitation float64 `json:"precipitation"`
	TempMax       float64 `json:"temp_max"`
	TempMin       float64 `json:"temp_min"`
	Wind          float64 `json:"wind"`
	Weather       string  `json:"weather"`
	Corrupted     bool    `json:"-"`
}
