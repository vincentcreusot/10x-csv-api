package main

import (
	"fmt"
	"os"

	"github.com/vincentcreusot/10x-csv-api/pkg/api"
	"github.com/vincentcreusot/10x-csv-api/pkg/csv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please specify the file to parse")
		os.Exit(0)
	}
	filepath := os.Args[1]

	fileLines := csv.ParseCsv(filepath)

	wlapi := api.NewWeatherApi(fileLines)
	router := wlapi.SetupRouter()
	wlapi.Start(router)

}
