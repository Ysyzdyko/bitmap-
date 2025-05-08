package filter

import (
	strct "bitmap/structure"
)

func FilterBlue(d *strct.BMPdata) {
	for _, row := range d.Color {
		for _, pixel := range row {
			pixel.Green = 0
			pixel.Red = 0
		}
	}
}
