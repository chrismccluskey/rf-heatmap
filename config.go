package main

import (
	"flag"
	"fmt"
	"strconv"
)

type Config struct {
	offset int
	csv_filename string
	png_filename string
	zoom float64
}

func (c *Config) Build() {
	offset := flag.Int("offset", 0, "offset in hz")
	csv_filename := flag.String("csv", "", "input csv filename")
	png_filename := flag.String("png", "", "output png filename")
	zoom := flag.String("zoom", "1.0", "frequency zoom")

	flag.Parse()

	c.offset = *offset
	c.csv_filename = *csv_filename
	c.png_filename = *png_filename
	c.zoom, _ = strconv.ParseFloat(*zoom, 64)
}

func (c Config) Print() {
	fmt.Println("Configuration:")

	fmt.Printf("Frequency Offset:	%d Hz\n", c.offset)
	fmt.Printf("Input CSV File:		%s\n", c.csv_filename)
	fmt.Printf("Output PNG File:	%s\n", c.png_filename)
	fmt.Printf("Zoom:				%f\n", c.zoom)
}
