package utils

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	strct "bitmap/structure"
)

func ReadBMPHeader(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	bmpHeader := strct.BMPHeader{}
	if err := binary.Read(file, binary.LittleEndian, &bmpHeader); err != nil {
		return fmt.Errorf("failed to read BMP header: %v", err)
	}

	if string(bmpHeader.Signature[:]) != "BM" {
		return fmt.Errorf("not a valid BMP file")
	}

	dibHeader := strct.DIBHeader{}
	if err := binary.Read(file, binary.LittleEndian, &dibHeader); err != nil {
		return fmt.Errorf("failed to read DIB header: %v", err)
	}

	var builder strings.Builder
	fmt.Println("BMP Header:")
	fmt.Printf("- FileType: %s\n", bmpHeader.Signature)
	fmt.Printf("- FileSizeInBytes: %d\n", bmpHeader.FileSize)
	fmt.Printf("- HeaderSize: %d\n", bmpHeader.DataOffset)
	fmt.Println("DIB Header:")
	fmt.Printf("- DibHeaderSize: %d\n", dibHeader.Size)
	fmt.Printf("- WidthInPixels: %d\n", dibHeader.Width)
	fmt.Printf("- HeightInPixels: %d\n", dibHeader.Height)
	fmt.Printf("- PixelSizeInBits: %d\n", dibHeader.BitsPerPixel)
	fmt.Printf("- ImageSizeInBytes: %d\n", dibHeader.ImageSize)
	fmt.Print(builder.String())

	return nil
}
