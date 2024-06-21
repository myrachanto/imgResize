package imgResize

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

// Imagetype processes and resizes an image based on its type.
func ResizeImage(f, filename string, height, width int) error {
	file, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("error opening the file")
	}
	defer file.Close()

	if imageSizeChecker(file) {
		log.Println("image size is smaller than 100kbs")
		return nil
	}

	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return fmt.Errorf("error reading the file")
	}

	filetype := http.DetectContentType(buff)
	newWidth := uint(width)
	newHeight := uint(0) // Automatically calculate to maintain aspect ratio

	switch filetype {
	case "image/jpeg", "image/jpg":
		if err := resizeJPG(file, filename, newHeight, newWidth); err != nil {
			return err
		}
	case "image/png":
		if err := resizePng(file, filename, newHeight, newWidth); err != nil {
			return err
		}
	case "image/webp":
		if err := resizeWebP(file, filename, newHeight, newWidth); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported file format")
	}

	return nil
}

// resizePng resizes a PNG image.
func resizePng(file *os.File, filename string, height, width uint) error {
	img, err := png.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding PNG")
	}

	m := resize.Resize(width, height, img, resize.Lanczos3)
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating PNG file")
	}
	defer out.Close()

	if err := png.Encode(out, m); err != nil {
		return fmt.Errorf("error encoding PNG")
	}
	return nil
}

// resizeJPG resizes a JPEG image.
func resizeJPG(file *os.File, filename string, height, width uint) error {
	img, err := jpeg.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding JPEG")
	}

	m := resize.Resize(width, height, img, resize.Lanczos3)
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating JPEG file")
	}
	defer out.Close()

	if err := jpeg.Encode(out, m, nil); err != nil {
		return fmt.Errorf("error encoding JPEG")
	}
	return nil
}

// resizeWebP resizes a WebP image.
func resizeWebP(file *os.File, filename string, height, width uint) error {
	img, err := webp.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding WebP")
	}

	m := resize.Resize(width, height, img, resize.Lanczos3)
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating WebP file")
	}
	defer out.Close()

	if err := webp.Encode(out, m, nil); err != nil {
		return fmt.Errorf("error encoding WebP")
	}
	return nil
}

// imageSizeChecker checks if the file size is less than 100KB.
func imageSizeChecker(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("error getting file info: %v", err)
		return false
	}
	const maxSize int64 = 100 * 1024
	return fileInfo.Size() < maxSize
}
