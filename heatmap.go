package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type heatmap struct {
	readings []reading
	config Config
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

	f, err := os.Create(h.config.png_filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

