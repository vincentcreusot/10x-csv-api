package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincentcreusot/10x-csv-api/pkg/structs"
)

type WeatherApi struct {
	Lines []structs.WeatherLine
}

func NewWeatherApi(lines []structs.WeatherLine) *WeatherApi {
	return &WeatherApi{
		Lines: lines,
	}
}

func (a *WeatherApi) SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/query", a.getWeatherLines)

	return r
}

func (a *WeatherApi) getWeatherLines(c *gin.Context) {
	// Get date filter query param
	dateFilter := c.Query("date")
	limit := c.Query("limit")
	weatherFilter := c.Query("weather")

	// Filter lines by date if provided
	var resultLines []structs.WeatherLine
	if dateFilter != "" {
		for _, line := range a.Lines {
			if line.Date.Format("2006-01-02") == dateFilter {
				resultLines = append(resultLines, line)
			}
		}
	} else {
		// No date filter, so start with all lines
		resultLines = a.Lines
	}

	// Filter by weather if provided
	if weatherFilter != "" {
		var weatherFiltered []structs.WeatherLine
		for _, line := range resultLines {
			if line.Weather == weatherFilter {
				weatherFiltered = append(weatherFiltered, line)
			}
		}
		resultLines = weatherFiltered
	}

	// Apply limit filter
	if limit != "" {
		limitNum, _ := strconv.Atoi(limit)
		resultLines = resultLines[:limitNum]
	}

	c.JSON(200, resultLines)
}
