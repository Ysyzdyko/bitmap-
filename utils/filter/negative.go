package filter

import (
	strct "bitmap/structure"
)

func FilterNegative(d *strct.BMPdata) {
	for _, row := range d.Color {
		for _, pixel := range row {
			pixel.Red = 255 - pixel.Red
			pixel.Green = 255 - pixel.Green
			pixel.Blue = 255 - pixel.Blue
		}
	}
}
