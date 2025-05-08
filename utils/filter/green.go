package filter

import (
	strct "bitmap/structure"
)

func FilterGreen(d *strct.BMPdata) {
	for _, row := range d.Color {
		for _, pixel := range row {
			pixel.Red = 0
			pixel.Blue = 0
		}
	}
}
