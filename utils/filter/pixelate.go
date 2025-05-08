package filter

import (
	strct "bitmap/structure"
)

var pixelateLevel int = 2

func FilterPixelate(d *strct.BMPdata) {
	height := len(d.Color)
	width := len(d.Color[0])

	blockSize := (height / 10) * pixelateLevel // Calculate block size based on image height
	if width/10 < blockSize {
		blockSize = (width / 10) * pixelateLevel // Adjust block size if image width is smaller
	}
	if blockSize == 0 {
		blockSize = 1
	}

	for i := 0; i < height; i += blockSize {
		for j := 0; j < width; j += blockSize {
			var sumRed, sumGreen, sumBlue int
			count := 0

			for bi := i; bi < i+blockSize && bi < height; bi++ {
				for bj := j; bj < j+blockSize && bj < width; bj++ {
					pixel := d.Color[bi][bj]
					sumRed += int(pixel.Red)
					sumGreen += int(pixel.Green)
					sumBlue += int(pixel.Blue)
					count++
				}
			}

			// Calculate the average color for the block
			if count > 0 {
				avgRed := uint8(sumRed / count)
				avgGreen := uint8(sumGreen / count)
				avgBlue := uint8(sumBlue / count)

				for bi := i; bi < i+blockSize && bi < height; bi++ {
					for bj := j; bj < j+blockSize && bj < width; bj++ {
						pixel := d.Color[bi][bj]
						pixel.Red = avgRed
						pixel.Green = avgGreen
						pixel.Blue = avgBlue
					}
				}
			}
		}
	}
	pixelateLevel++
}
