package utils

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	strct "bitmap/structure"
)

func SaveBMP(d *strct.BMPdata, filename string) error {
	outputPath := filename

	if strings.Contains(outputPath, "output_images") {
		outputFileName := filepath.Base(filename)
		if outputFileName == filepath.Base(filename) {
			outputPath = outputFileName
		}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := binary.Write(file, binary.LittleEndian, d.BMP); err != nil {
		return err
	}

	if err := binary.Write(file, binary.LittleEndian, d.DIB); err != nil {
		return err
	}
	// Seek to the beginning of the pixel data
	if _, err := file.Seek(int64(d.BMP.DataOffset), io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek to pixel data: %v", err)
	}
	// Write the pixel data
	for i := int32(d.DIB.Height) - 1; i >= 0; i-- {
		for j := int32(0); j < d.DIB.Width; j++ {
			if err := binary.Write(file, binary.LittleEndian, d.Color[i][j]); err != nil {
				return fmt.Errorf("failed to write pixel data: %v", err)
			}
		}
		// Add padding to align the rows to a 4-byte boundary
		padding := (4 - (d.DIB.Width*3)%4) % 4
		if padding > 0 {
			if _, err := file.Write(make([]byte, padding)); err != nil {
				return err
			}
		}
	}
	// Log success message after saving
	fmt.Printf("Image saved to %s\n", outputPath)
	return nil
}
