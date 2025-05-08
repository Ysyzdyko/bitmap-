package structure

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

type BMPHeader struct {
	Signature  [2]byte
	FileSize   uint32
	Reserved   [4]byte
	DataOffset uint32
}

type DIBHeader struct {
	Size            uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XpixelsPerM     int32
	YpixelsPerM     int32
	ColorsUsed      uint32
	ImportantColors uint32
}

// Pixel represents a single pixel in the image
type Pixel struct {
	Blue  uint8
	Green uint8
	Red   uint8
}

type BMPdata struct {
	BMP   *BMPHeader
	DIB   *DIBHeader
	Color [][]*Pixel
}

// Method to read pixel data from a reader
func (p *Pixel) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, p)
}

func ReadBMP(filename string) (*BMPdata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	bmpHeader := &BMPHeader{}
	if err := binary.Read(file, binary.LittleEndian, bmpHeader); err != nil {
		return nil, fmt.Errorf("failed to read BMP header: %v", err)
	}

	if string(bmpHeader.Signature[:]) != "BM" {
		return nil, fmt.Errorf("not a valid BMP file")
	}

	// Read the DIB header
	dibHeader := &DIBHeader{}
	if err := binary.Read(file, binary.LittleEndian, dibHeader); err != nil {
		return nil, fmt.Errorf("failed to read DIB header: %v", err)
	}
	if dibHeader.BitsPerPixel != 24 {
		log.Fatal("not a 24-bit color bitmap file")
	}
	// Seek to the start of the pixel data
	if _, err := file.Seek(int64(bmpHeader.DataOffset), io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek to pixel data: %v", err)
	}

	color := make([][]*Pixel, dibHeader.Height)
	for i := range color {
		color[i] = make([]*Pixel, dibHeader.Width)
	}

	for i := int(dibHeader.Height) - 1; i >= 0; i-- {
		for j := 0; j < int(dibHeader.Width); j++ {
			pixel := &Pixel{}
			if err := pixel.Read(file); err != nil {
				return nil, fmt.Errorf("failed to read pixel data: %v", err)
			}
			color[i][j] = pixel
		}
		// Handle padding (BMP rows are aligned to 4-byte boundaries)
		padding := (4 - (dibHeader.Width*3)%4) % 4
		if padding > 0 {
			if _, err := file.Seek(int64(padding), io.SeekCurrent); err != nil {
				return nil, err
			}
		}
	}

	return &BMPdata{
		BMP:   bmpHeader,
		DIB:   dibHeader,
		Color: color,
	}, nil
}
