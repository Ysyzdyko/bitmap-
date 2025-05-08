package utils

import (
	"fmt"
	"strconv"
	"strings"

	strct "bitmap/structure"
)

func CropImage(d *strct.BMPdata, crops []string) error {
	// Iterate over each crop parameter set
	for _, cropParams := range crops {
		err := applyCrop(d, cropParams)
		if err != nil {
			return err
		}
	}
	return nil
}

func applyCrop(d *strct.BMPdata, cropParams string) error {
	// Split the crop parameters into components
	params := strings.Split(cropParams, "-")
	if len(params) != 2 && len(params) != 4 {
		return fmt.Errorf("invalid crop parameters: %s", cropParams)
	}

	var err error
	var offsetX, offsetY, width, height int

	offsetX, err = strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("invalid offset X: %v", err)
	}

	offsetY, err = strconv.Atoi(params[1])
	if err != nil {
		return fmt.Errorf("invalid offset Y: %v", err)
	}

	// If width and height are specified
	if len(params) == 4 {
		width, err = strconv.Atoi(params[2])
		if err != nil {
			return fmt.Errorf("invalid width: %v", err)
		}

		height, err = strconv.Atoi(params[3])
		if err != nil {
			return fmt.Errorf("invalid height: %v", err)
		}
	} else {
		// Use the remaining distance to the edge of the image
		width = int(d.DIB.Width) - offsetX
		height = int(d.DIB.Height) - offsetY
	}

	if offsetX < 0 || offsetY < 0 || width <= 0 || height <= 0 {
		return fmt.Errorf("invalid crop dimensions")
	}

	if offsetX+width > int(d.DIB.Width) || offsetY+height > int(d.DIB.Height) {
		return fmt.Errorf("crop dimensions exceed image size")
	}

	// Perform the crop operation
	newColor := make([][]*strct.Pixel, height)
	for i := 0; i < height; i++ {
		newColor[i] = make([]*strct.Pixel, width)
		copy(newColor[i], d.Color[offsetY+i][offsetX:offsetX+width])
	}

	// Update the BMP data with the new cropped image
	d.Color = newColor
	d.DIB.Width = int32(width)
	d.DIB.Height = int32(height)
	d.DIB.ImageSize = uint32(height * ((width*3 + 3) &^ 3))
	d.BMP.FileSize = d.BMP.DataOffset + d.DIB.ImageSize

	return nil
}
