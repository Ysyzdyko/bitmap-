# Bitmap Team Project

Welcome to the Bitmap team project for the Foundation Course at Alem School!

### Team Members
- **Team Leader:** ssaidaly
- **Members:** nskakova, basagitz, ysyzdyko, dabektur

### Overview
This program is designed to work with bitmap image files in Go. 

### Building the Program
To build the program, run the following command:

```bash
go build -o bitmap .
```

### Main Commands
The program has two primary commands: **header** and **apply**.

#### 1. Header
The **header** command prints information about a BMP file.

**Usage:**
```bash
./bitmap header sample.bmp
```

#### 2. Apply
The **apply** command allows you to modify BMP files with several options.

**Usage:**
```bash
./bitmap apply [options] <source_file> [<source_file> ...]
```

### Apply Subcommands
Here are the available subcommands for **apply**:

#### a. Mirror
Mirrors the BMP image.

**Usage:**
```bash
./bitmap apply --mirror=[h, hor, horizontal, horizontally, v, ver, vertical, vertically] <source_file> <result_file>
```

#### b. Filter
Filters a chosen color or applies a specific filter to the BMP image.

**Usage:**
```bash
./bitmap apply --filter=[blue, red, green, negative, pixelate, grayscale, blur] <source_file> <result_file>
```

#### c. Crop
Crops the BMP image to the specified dimensions.

**Usage:**
```bash
./bitmap apply --crop=[width,height] <source_file> <result_file>
```

#### d. Rotate
Rotates the BMP image left or right by specified degrees.

**Usage:**
```bash
./bitmap apply --rotate=[right, 90, +90, 180, +180, 270, +270, left, -90, -180, -270] <source_file> <result_file>
```

### Help Command
You can always access the help command for guidance on using the program and its subcommands.

**Usage:**
```bash
./bitmap --help
./bitmap header --help
./bitmap apply --help
```

Feel free to reach out with any questions or for assistance! Enjoy working with the Bitmap program!