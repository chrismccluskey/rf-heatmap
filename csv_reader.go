package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"strconv"
	"strings"
)

type CSVReader struct {}
type DataSet struct {}

type reading struct {
	date string
	time string
	hz_low float64
	hz_high float64
	hz_step float64
	samples int
	dbs []float64
}

func (d *CSVReader) Read() (filename string) {
	// statistics of data file
	peak_db := float64(-1000)
	valley_db := float64(0)
	lowest_hz := float64(999999999999)
	highest_hz := float64(0)
	readings := []reading{}
	hz_width := 1.0
	hz_step := 1.0
	rows := 0
	// read
	file, err := os.Open(filename)
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
}
