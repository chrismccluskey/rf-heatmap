package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"io"
	"strconv"
	"strings"
)

type reading struct {
	date string
	time string
	hz_low float64
	hz_high float64
	hz_step float64
	samples int
	dbs []float64
}

type heatmap struct {
	readings []reading
	peak_db float64
	valley_db float64
	db_difference float64
	db_multiplier float64
	lowest_hz float64
	highest_hz float64
	hz_difference float64
	hz_multiplier float64
	hz_width float64
	rows int
}

func (h *heatmap) draw() {
	var img = image.NewRGBA(image.Rect(0, 0, int(h.hz_width), h.rows))
	var col color.Color

	// draw
	x := -1
	y := -1
	for _, reading := range h.readings {
		if ( reading.hz_low == h.lowest_hz ) {
			x = -1
			y = y+1
		}
		for _, db := range reading.dbs {
			x = x+1
			heat := uint8((db-h.valley_db)*h.db_multiplier)
			col = color.RGBA{0, heat, heat, 255}
			img.Set(x, y, col)
		}
	}

	f, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func main() {
	//offset := float64(-125000000)
	peak_db := float64(-1000)
	valley_db := float64(0)
	lowest_hz := float64(999999999999)
	highest_hz := float64(0)
	readings := []reading{}
	hz_multiplier := 1.0
	hz_width := 1.0
	hz_step := 1.0
	rows := 0

	// read
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Comma = ','
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		date := record[0]
		time := record[1]
		hz_low, err := strconv.ParseFloat(strings.Trim(record[2], " "), 64)
		hz_high, err := strconv.ParseFloat(strings.Trim(record[3], " "), 64)
		hz_step, err = strconv.ParseFloat(strings.Trim(record[4], " "), 64)
		samples, err := strconv.Atoi(record[5])
		//fmt.Printf("%s %s : %f - %f\n", date, time, hz_low, hz_high)

		if ( hz_low < lowest_hz ) {
			lowest_hz = hz_low
		}
		if ( hz_high > highest_hz ) {
			highest_hz = hz_high
		}
		if ( hz_low == lowest_hz ) {
			rows = rows + 1
		}

		
		dbs := []float64{}
		for i := 6; i < len(record); i++ {
			//n := float64(i) - 5.0;
			db, err := strconv.ParseFloat(strings.Trim(record[i], " "), 64)
			if ( db > peak_db ) {
				peak_db = db
			}
			if ( db < valley_db ) {
				valley_db = db
			}
			if ( err != nil ) {
				fmt.Println("Error: ", err)
			}
			//fmt.Printf("%f: %f\n", (hz_low+n*hz_step)+offset, db)
			dbs = append(dbs, db)
		}
		r := reading{date: date, time: time, hz_low: hz_low, hz_high: hz_high, hz_step: hz_step, samples: samples, dbs: dbs}
		readings = append(readings, r)
	}

	db_difference := peak_db - valley_db
	db_multiplier := 255.0 / db_difference
	hz_difference := highest_hz - lowest_hz
	hz_width = hz_difference / hz_step
	fmt.Printf("Peak Db: %f	Valley Db: %f	Difference: %f	Multiplier: %f	Hz Width: %f\n", peak_db, valley_db, db_difference, db_multiplier, hz_width)
	h := heatmap{readings: readings, peak_db: peak_db, valley_db: valley_db, db_difference: db_difference, db_multiplier: db_multiplier, lowest_hz: lowest_hz, highest_hz: highest_hz, hz_difference: hz_difference, hz_multiplier: hz_multiplier, hz_width: hz_width, rows: rows}
	h.draw()
}
