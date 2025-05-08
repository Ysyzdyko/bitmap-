package utils

import (
	strct "bitmap/structure"
)

func MirrorHorizontal(d *strct.BMPdata) {
	for i := range d.Color {
		for j := 0; j < len(d.Color[i])/2; j++ {
			d.Color[i][j], d.Color[i][len(d.Color[i])-1-j] = d.Color[i][len(d.Color[i])-1-j], d.Color[i][j]
		}
	}
}

func MirrorVertical(d *strct.BMPdata) {
	for i := 0; i < len(d.Color)/2; i++ {
		d.Color[i], d.Color[len(d.Color)-1-i] = d.Color[len(d.Color)-1-i], d.Color[i]
	}
}
