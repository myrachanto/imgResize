package imgResize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResizeImage_ValidJPEG(t *testing.T) {
	err := ResizeImage("./testdata/test.jpg", "output/test_resized.jpg", 200, 300)
	assert.NoError(t, err, "Expected no error for valid JPEG")
	// Add assertions to validate output dimensions and file creation
}

func TestResizeImage_ValidPNG(t *testing.T) {
	err := ResizeImage("testdata/test.png", "output/test_resized.png", 300, 200)
	assert.NoError(t, err, "Expected no error for valid PNG")
}

// func TestResizeImage_ValidWebP(t *testing.T) {
// 	err := ResizeImage("testdata/test.webp", "output/test_resized.webp", 300, 300)
// 	assert.NoError(t, err, "Expected no error for valid WebP")
// }


func TestResizeImage_FileNotFound(t *testing.T) {
	err := ResizeImage("testdata/nonexistent.jpg", "output/test_resized.jpg", 200, 200)
	assert.Error(t, err, "Expected error for nonexistent file")
}

func TestResizeImage_FileTooSmall(t *testing.T) {
	err := ResizeImage("testdata/small.png", "output/test_resized.png", 200, 200)
	assert.NoError(t, err, "Expected no error but no resizing for small file")
}

func TestResizeImage_ResizeToSmallDimensions(t *testing.T) {
	err := ResizeImage("testdata/test.jpg", "output/test_resized.jpg", 1, 1)
	assert.NoError(t, err, "Expected no error for resizing to small dimensions")
}

func TestResizeImage_InvalidOutputPath(t *testing.T) {
	err := ResizeImage("testdata/test.jpg", "/invalid/path/test_resized.jpg", 200, 200)
	assert.Error(t, err, "Expected error for invalid output path")
}

func TestResizeImage_CorruptedImage(t *testing.T) {
	err := ResizeImage("testdata/corrupted.jpg", "output/test_resized.jpg", 200, 200)
	assert.Error(t, err, "Expected error for corrupted image")
}
