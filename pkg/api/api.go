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

func (a *WeatherApi) Start(router *gin.Engine) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown: ", err)
	}

	logger.Info("Server exiting")
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
	if limit != "" {
		limitNum, _ := strconv.Atoi(limit)
		limitApplied := int(math.Min(float64(len(resultLines)), float64(limitNum)))
		resultLines = resultLines[:limitApplied]
	}

	c.JSON(200, resultLines)
}
