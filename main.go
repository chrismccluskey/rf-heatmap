package main

import (
	"fmt"
)

func main() {
	// configurable settings
	config := Config{}
	config.Build()
	config.Print()

	c := CSVReader{}
	data_set := c.Read(config.csv_filename)

	fmt.Printf("%s", data_set)

	h := heatmap{data_set, config}
	h.draw()
}
