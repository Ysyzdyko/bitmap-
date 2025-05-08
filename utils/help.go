package utils

func MainHelp() string {
	return `Usage:
  bitmap <command> [arguments]

The commands are:
  header    prints bitmap file header information
  apply     applies processing to the image and saves it to the file`
}

func HelpHeader() string {
	return `Usage:
  bitmap header <source_file>

Description:
  Prints bitmap file header information`
}

func HelpApply() string {
	return `Usage:
  bitmap apply [options] <source_file> [<source_file> ...]

Options:
  -h, --help       prints program usage information
  --crop           trims a bitmap image according to specified parameters
                   accepts either 2 or 4 values in pixels:
                   --crop=OffsetX-OffsetY[-Width-Height]
                   can be specified multiple times
  --rotate         rotates a bitmap image by a specified angle
                   options: right, left, 90, -90, 180, -180, 270, -270
                   can be specified multiple times
  --filter         filter to apply (can be specified multiple times)
                   e.g., --filter=blur --filter=grayscale
  --mirror         mirrors the image (horizontal or vertical)
  --output         directory to save processed images (default: output_images)`
}
