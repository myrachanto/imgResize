# imgResize Package
The imgResize package is a Golang package that provides a simple way to resize image files in various formats (JPEG, PNG, and WebP). It includes functionality to check the size of the image and skip resizing if the file is smaller than 100KB. The package also handles resizing of images while maintaining their aspect ratio and allows you to save the resized images in the specified format.

### Features
- Resizes images in JPEG, PNG, and WebP formats.
- Automatically maintains aspect ratio while resizing.
- Skips resizing if the image size is smaller than 100KB.
- Handles errors related to file reading, decoding, resizing, and encoding.

### Installation

To use the imgResize package in your Go project, you can import it using the following:

To import the package
```bash
go get github.com:myrachanto/imgResize.git
```

### Resize an Image

You can use the ResizeImage function to resize an image. It will automatically detect the image format (JPEG, PNG) and resize it accordingly.

Example

```bash
package main

import (
	"log"
	"github.com:myrachanto/imgResize"
)

func main() {
	// Resize an image to 800x600
	err := imgResize.ResizeImage("input.jpg", "output.jpg", 600, 800)
	if err != nil {
		log.Fatalf("Error resizing image: %v", err)
	}
}
```

- func ResizeImage(f, filename string, height, width int) error
- f: Path to the input image file.
- filename: Path to the output resized image.
- height: Desired height for the resized image (0 for automatic aspect ratio calculation).
- width: Desired width for the resized image.

### Supported Formats
The package supports the following image formats:

- JPEG (image/jpeg, image/jpg)
- PNG (image/png)

The package will automatically skip resizing images that are smaller than 100KB in size. You can check for this behavior by logging the message: image size is smaller than 100kbs.
