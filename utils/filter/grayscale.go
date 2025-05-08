package filter

import (
	strct "bitmap/structure"
)

func FilterGrayScale(d *strct.BMPdata) {
	for _, row := range d.Color {
		for _, pixel := range row {
			gray := uint8(0.3*float64(pixel.Red) + 0.59*float64(pixel.Green) + 0.11*float64(pixel.Blue))
			pixel.Red = gray
			pixel.Green = gray
			pixel.Blue = gray
		}
	}
}
