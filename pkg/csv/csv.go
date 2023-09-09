package csv

import (
	"encoding/csv"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/vincentcreusot/10x-csv-api/pkg/structs"
)

var logger *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
var dateLayout = "2006-01-02"

func ParseCsv(filepath string) []structs.WeatherLine {

	// open file
	f, err := os.Open(filepath)
	if err != nil {
		logger.Error("Error opening csv file", err)
		os.Exit(2)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		logger.Error("Error reading file", err)
	}

	return parseLines(data)
}

func parseLines(parsedData [][]string) []structs.WeatherLine {
	lines := make([]structs.WeatherLine, 0)
	for i, line := range parsedData {
		if i > 0 {
			wl := structs.WeatherLine{}
			var err error
			wl.Date, err = time.Parse(dateLayout, line[0])
			if err != nil {
				logger.Info("error parsing ̈field date for value "+line[0], err)
				wl.Corrupted = true
			}
			wl.Precipitation = parseNumber(line[1], "precipitation", &wl)
			wl.TempMax = parseNumber(line[2], "temp_max", &wl)
			wl.TempMin = parseNumber(line[3], "temp_min", &wl)
			wl.Wind = parseNumber(line[4], "wind", &wl)
			wl.Weather = line[5]
			if !wl.Corrupted {
				lines = append(lines, wl)
			}
		}

	}
	return lines
}

func parseNumber(value string, field string, wo *structs.WeatherLine) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		logger.Info("error parsing ̈field "+field+"for value "+value, err)
	}
	return f
}
