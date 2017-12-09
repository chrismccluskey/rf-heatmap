package main

import (
//	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type heatmap struct {
	data_set DataSet
	config Config
}

func (h *heatmap) draw() {
	var img = image.NewRGBA(image.Rect(0, 0, int(h.data_set.hz_width), h.data_set.rows))
	var col color.Color

	// draw
	x := -1
	y := -1
	for _, reading := range h.data_set.readings {
		if ( reading.hz_low == h.data_set.lowest_hz ) {
			x = -1
			y = y+1
		}
		for _, db := range reading.dbs {
			x = x+1
			heat := uint8((db-h.data_set.valley_db)*h.data_set.db_multiplier)
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

