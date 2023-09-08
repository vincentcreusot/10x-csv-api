package csv

import (
	"encoding/csv"
	"log/slog"
	"os"

	"github.com/vincentcreusot/10x-csv-api/pkg/structs"
)

var logger *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func ParseCsv(filepath string) []structs.WeatherLine {

	// open file
	f, err := os.Open("data.csv")
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

	lines := make([]structs.WeatherLine, 0)
	for i, line := range data {
		if i > 0 {
			wl := structs.WeatherLine{}
			wl.Date = line[0]

			lines = append(lines, wl)
		}

	}
	return lines
}
