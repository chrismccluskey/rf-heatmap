package main

import (
	"fmt"
)

func main() {
	// configurable settings
	config := Config{}
	config.Build()
	config.Print()

	fmt.Printf("Peak Db: %f	Valley Db: %f	Difference: %f	Multiplier: %f	Hz Width: %f\n", peak_db, valley_db, db_difference, db_multiplier, hz_width)

	h := heatmap{config: config, readings: readings, peak_db: peak_db, valley_db: valley_db, db_difference: db_difference, db_multiplier: db_multiplier, lowest_hz: lowest_hz, highest_hz: highest_hz, hz_difference: hz_difference, hz_multiplier: hz_multiplier, hz_width: hz_width, rows: rows}
	h.draw()
}
