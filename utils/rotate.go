package utils

import (
	"fmt"
	"strconv"
	"strings"

	strct "bitmap/structure"
)

func RotateImage(d *strct.BMPdata, rotations []string) error {
	totalRotation := 0

	// Sum up all rotation angles
	for _, rot := range rotations {
		angle, err := parseRotation(rot)
		if err != nil {
			return err
		}
		totalRotation += angle
	}

	totalRotation = ((totalRotation % 360) + 360) % 360

	// Map of rotation functions based on the normalized angle
	rotationFuncs := map[int]func(*strct.BMPdata){
		90:  Rotate90,
		180: Rotate180,
		270: Rotate270,
		0:   func(d *strct.BMPdata) {},
	}

	// Get the rotation function based on the total rotation
	rotationFunc, exists := rotationFuncs[totalRotation]
	if !exists {
		return fmt.Errorf("unsupported rotation angle: %d", totalRotation)
	}

	// Apply the rotation function
	rotationFunc(d)
	return nil
}

func parseRotation(rotation string) (int, error) {
	rotationMap := map[string]int{
		"right": 90,
		"90":    90,
		"+90":   90,
		"left":  -90,
		"-90":   -90,
		"180":   180,
		"+180":  180,
		"-180":  180,
		"270":   270,
		"+270":  270,
		"-270":  -90,
	}

	rotation = strings.ToLower(rotation)
	if angle, exists := rotationMap[rotation]; exists {
		return angle, nil
	}

	// Attempt to parse rotation as an integer
	angle, err := strconv.Atoi(rotation)
	if err != nil {
		return 0, fmt.Errorf("invalid rotation value: %s", rotation)
	}
	return angle, nil
}

func Rotate90(d *strct.BMPdata) {
	width := int(d.DIB.Width)
	height := int(d.DIB.Height)

	newColor := make([][]*strct.Pixel, width)
	for i := range newColor {
		newColor[i] = make([]*strct.Pixel, height)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newColor[j][height-1-i] = d.Color[i][j]
		}
	}

	// Update the image data with the rotated image
	d.Color = newColor
	d.DIB.Width, d.DIB.Height = int32(height), int32(width)
}

func Rotate180(d *strct.BMPdata) {
	width := int(d.DIB.Width)
	height := int(d.DIB.Height)

	for i := 0; i < height/2; i++ {
		d.Color[i], d.Color[height-1-i] = d.Color[height-1-i], d.Color[i]
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width/2; j++ {
			d.Color[i][j], d.Color[i][width-1-j] = d.Color[i][width-1-j], d.Color[i][j]
		}
	}
}

func Rotate270(d *strct.BMPdata) {
	width := int(d.DIB.Width)
	height := int(d.DIB.Height)

	newColor := make([][]*strct.Pixel, width)
	for i := range newColor {
		newColor[i] = make([]*strct.Pixel, height)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newColor[width-1-j][i] = d.Color[i][j]
		}
	}

	d.Color = newColor
	d.DIB.Width, d.DIB.Height = int32(height), int32(width)
}
