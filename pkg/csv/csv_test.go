package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vincentcreusot/10x-csv-api/pkg/structs"
)

func TestParseLines(t *testing.T) {

	t.Run("Parses lines correctly", func(t *testing.T) {
		input := [][]string{
			{"date", "precipitation", "temp_max", "temp_min", "wind", "weather"},
			{"2012-01-01", "0.0", "12.8", "5.0", "4.7", "drizzle"},
		}

		expected := []structs.WeatherLine{
			{
				Date:          "2012-01-01",
				Precipitation: 0.0,
				TempMax:       12.8,
				TempMin:       5.0,
				Wind:          4.7,
				Weather:       "drizzle",
			},
		}

		output := parseLines(input)

		assert.Equal(t, expected, output)
	})

	t.Run("Skips header row", func(t *testing.T) {
		input := [][]string{
			{"date", "precipitation", "temp_max", "temp_min", "wind", "weather"},
			{"Value1", "Value2", "Value3", "Value4", "Value5", "Value6"},
		}

		assert.Len(t, parseLines(input), 1)
	})
}
