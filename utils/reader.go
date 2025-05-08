package utils

import (
	"encoding/binary"
	"fmt"
	"os"

	strct "bitmap/structure"
)

func ReadHeaders(filename string) (*strct.BMPHeader, *strct.DIBHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	bmpHeader := &strct.BMPHeader{}
	if err := binary.Read(file, binary.LittleEndian, bmpHeader); err != nil {
		return nil, nil, fmt.Errorf("failed to read BMP header: %v", err)
	}

	if string(bmpHeader.Signature[:]) != "BM" {
		return nil, nil, fmt.Errorf("not a valid BMP file")
	}

	dibHeader := &strct.DIBHeader{}
	if err := binary.Read(file, binary.LittleEndian, dibHeader); err != nil {
		return nil, nil, fmt.Errorf("failed to read DIB header: %v", err)
	}

	return bmpHeader, dibHeader, nil
}
