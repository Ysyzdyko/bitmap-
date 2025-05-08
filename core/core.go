package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	strct "bitmap/structure"
	u "bitmap/utils"
	f "bitmap/utils/filter"
)

type Filters []string

// Method for Filters
func (f *Filters) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func (f *Filters) String() string {
	return strings.Join(*f, ",")
}

type Rotations []string

// Method for Rotations
func (r *Rotations) Set(value string) error {
	*r = append(*r, value)
	return nil
}

func (r *Rotations) String() string {
	return strings.Join(*r, ",")
}

type Crops []string

// Method for Crops
func (c *Crops) Set(value string) error {
	*c = append(*c, value)
	return nil
}

func (c *Crops) String() string {
	return strings.Join(*c, ",")
}

type Mirrors []string

// Method for Mirrors
func (m *Mirrors) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func (m *Mirrors) String() string {
	return strings.Join(*m, ",")
}

// Main application structure
type Application struct{}

func (app *Application) Run() error {
	if len(os.Args) < 2 {
		fmt.Println(u.MainHelp())
		return nil
	}

	commandMap := map[string]func([]string) error{
		"header": app.runHeader,
		"apply":  app.runApply,
	}

	cmd := os.Args[1]

	if commandFunc, exists := commandMap[cmd]; exists {
		return commandFunc(os.Args[2:])
	}

	fmt.Println(u.MainHelp())
	return nil
}

func (app *Application) runHeader(args []string) error {
	headerCmd := flag.NewFlagSet("header", flag.ExitOnError)
	help := headerCmd.Bool("help", false, "Prints program usage information")
	headerCmd.BoolVar(help, "h", false, "Prints program usage information")
	headerCmd.Parse(args)

	if *help || headerCmd.NArg() != 1 {
		fmt.Println(u.HelpHeader())
		return fmt.Errorf("just one file")
	}

	inputFile := headerCmd.Arg(0)
	err := u.ReadBMPHeader(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	return nil
}

func (app *Application) runApply(args []string) error {
	applyCmd := flag.NewFlagSet("apply", flag.ExitOnError)
	help := applyCmd.Bool("help", false, "Prints program usage information")
	applyCmd.BoolVar(help, "h", false, "Prints program usage information")
	var filters Filters
	applyCmd.Var(&filters, "filter", "Filter to apply (can be specified multiple times)")

	var rotations Rotations
	applyCmd.Var(&rotations, "rotate", "Rotate angle (can be specified multiple times)")

	var crops Crops
	applyCmd.Var(&crops, "crop", "Crop parameters (can be specified multiple times)")

	var mirrors Mirrors
	applyCmd.Var(&mirrors, "mirror", "Mirror type (can be specified multiple times)")

	applyCmd.Parse(args)

	if *help || applyCmd.NArg() == 0 {
		fmt.Println(u.HelpApply())
		return nil
	}

	if applyCmd.NArg() < 2 {
		fmt.Println("Usage: <input file> <output file>")
		return nil
	}

	outputDir := "output_images"
	// if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
	// 	log.Fatalf("Error creating output directory %s: %v", outputDir, err)
	// }

	inputFile := applyCmd.Arg(0)
	outputFileName := applyCmd.Arg(1)
	outputFile := filepath.Join(outputDir, outputFileName)

	d, err := strct.ReadBMP(inputFile)
	if err != nil {
		log.Printf("Error reading BMP %s: %v", inputFile, err)
		return err
	}

	for _, filterName := range filters {
		if err := app.applyFilter(filterName, d); err != nil {
			log.Printf("Error applying filter %s to %s: %v", filterName, inputFile, err)
			return err
		}
	}

	for _, mirrorType := range mirrors {
		normalizedMirror := normalizeMirrorType(mirrorType)
		if err := app.applyMirror(normalizedMirror, d); err != nil {
			log.Printf("Error applying mirror %s to %s: %v", normalizedMirror, inputFile, err)
			return err
		}
	}

	if len(rotations) > 0 {
		if err := u.RotateImage(d, rotations); err != nil {
			log.Printf("Error rotating image %s: %v", inputFile, err)
			return err
		}
	}

	if len(crops) > 0 {
		if err := u.CropImage(d, crops); err != nil {
			log.Printf("Error cropping image %s: %v", inputFile, err)
			return err
		}
	}

	log.Printf("Saving processed image to: %s", outputFile)
	if err := u.SaveBMP(d, outputFile); err != nil {
		log.Printf("Error saving BMP %s: %v", outputFile, err)
		return err
	}

	fmt.Printf("Processed image saved to %s\n", outputFile)
	return nil
}

// Function to normalize mirror types
func normalizeMirrorType(mirror string) string {
	switch strings.ToLower(mirror) {
	case "horizontal", "horizontally", "hor", "h":
		return "horizontal"
	case "vertical", "vertically", "ver", "v":
		return "vertical"
	default:
		return mirror
	}
}

// Method to apply filters
func (app *Application) applyFilter(filterName string, d *strct.BMPdata) error {
	filterMap := map[string]func(*strct.BMPdata){
		"blue":      f.FilterBlue,
		"red":       f.FilterRed,
		"green":     f.FilterGreen,
		"negative":  f.FilterNegative,
		"grayscale": f.FilterGrayScale,
		"pixelate":  f.FilterPixelate,
		"blur":      f.FilterBlur,
	}

	filterFunc, exists := filterMap[filterName]
	if !exists {
		return fmt.Errorf("unknown filter: %s", filterName)
	}

	filterFunc(d)
	return nil
}

// Method to apply mirror transformations
func (app *Application) applyMirror(mirrorType string, d *strct.BMPdata) error {
	mirrorMap := map[string]func(*strct.BMPdata){
		"horizontal": u.MirrorHorizontal,
		"vertical":   u.MirrorVertical,
	}

	mirrorFunc, exists := mirrorMap[mirrorType]
	if !exists {
		return fmt.Errorf("unknown mirror type: %s", mirrorType)
	}

	mirrorFunc(d)
	return nil
}
