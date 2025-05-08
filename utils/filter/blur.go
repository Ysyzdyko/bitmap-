package filter

import (
	strct "bitmap/structure"
)

var blurlevel = 1

func FilterBlur(d *strct.BMPdata) {
	height := len(d.Color)
	width := len(d.Color[0])

	kernelSize := blurlevel // Radius of the blur

	// Create a matrix to store the new pixel values
	newColor := make([][]*strct.Pixel, height)
	for i := range newColor {
		newColor[i] = make([]*strct.Pixel, width)
		for j := range newColor[i] {
			newColor[i][j] = &strct.Pixel{}
		}
	}

	// Apply the blur filter
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var sumRed, sumGreen, sumBlue int
			count := 0

			// Calculate the average color in the kernel
			for ki := -kernelSize; ki <= kernelSize; ki++ {
				for kj := -kernelSize; kj <= kernelSize; kj++ {
					ni := i + ki
					nj := j + kj
					if ni >= 0 && ni < height && nj >= 0 && nj < width {
						pixel := d.Color[ni][nj]
						sumRed += int(pixel.Red)
						sumGreen += int(pixel.Green)
						sumBlue += int(pixel.Blue)
						count++
					}
				}
			}

			// Set the new color as the average of the surrounding pixels
			if count > 0 {
				newColor[i][j].Red = uint8(sumRed / count)
				newColor[i][j].Green = uint8(sumGreen / count)
				newColor[i][j].Blue = uint8(sumBlue / count)
			}
		}
	}
	blurlevel++
	d.Color = newColor
}
