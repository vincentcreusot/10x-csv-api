package api

import (
	"context"
	"log/slog"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vincentcreusot/10x-csv-api/pkg/structs"
)

var logger *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// WeatherApi holder for the API
type WeatherApi struct {
	Lines []structs.WeatherLine
}

// NewWeatherApi constructs a new WeatherApi instance
// @param lines - The parsed CSV data
func NewWeatherApi(lines []structs.WeatherLine) *WeatherApi {
	return &WeatherApi{
		Lines: lines,
	}
}

// SetupRouter sets up the HTTP router and routes
func (a *WeatherApi) SetupRouter() *gin.Engine {
	r := gin.Default()

	// define /query as the path to get the weather lines
	r.GET("/query", a.getWeatherLines)

	return r
}

// Start starts the API server
func (a *WeatherApi) Start(router *gin.Engine) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown: ", err)
	}

	logger.Info("Server exiting")
}

func (a *WeatherApi) getWeatherLines(c *gin.Context) {
	dateFilter := c.Query("date")
	limitFilter := c.Query("limit")
	weatherFilter := c.Query("weather")

	// Filter lines by date if provided
	var resultLines []structs.WeatherLine
	if dateFilter != "" {
		for _, line := range a.Lines {
			if line.Date == dateFilter {
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
	if limitFilter != "" {
		limitNum, _ := strconv.Atoi(limitFilter)
		limitApplied := int(math.Min(float64(len(resultLines)), float64(limitNum)))
		resultLines = resultLines[:limitApplied]
	}

	c.JSON(200, resultLines)
}
