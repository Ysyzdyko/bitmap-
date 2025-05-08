package filter

import (
	strct "bitmap/structure"
)

func FilterRed(d *strct.BMPdata) {
	for _, row := range d.Color {
		for _, pixel := range row {
			pixel.Green = 0
			pixel.Blue = 0
		}
	}
}
