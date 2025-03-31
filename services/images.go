package services

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

func LoadImage(r *http.Request) (image.Image, string, error) {
	r.ParseMultipartForm(32 << 10)
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

// SaveImage saves an image to a file in the specified format (jpeg or png).
func SaveImage(img image.Image, format, outputPath string) error {
	// Create the output file
	outFile, err := os.Create(outputPath + "." + format)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode the image based on the format
	switch format {
	case "jpeg":
		err = jpeg.Encode(outFile, img, nil)
	case "png":
		err = png.Encode(outFile, img)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	return err
}

// This will be integrated with external storage bucket in future
func FileUpload(r *http.Request, filePath string) (string, error) {
	img, format, err := LoadImage(r)
	if err != nil {
		return "", err
	}

	// croppedImage, err := CropImage(img, 0, 0, 500, 500)
	// if err != nil {
	// 	return "", err
	// }
	resizedImg := ResizeImage(img, 2000, 1000)
	err = SaveImage(resizedImg, format, filePath+"resized_image_test")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Resized image saved to: %s", filePath), nil
}

func ResizeImage(img image.Image, width, height int) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Get the dimensions of the source image
	srcBounds := img.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	// Calculate scaling factors
	xScale := float64(srcWidth) / float64(width)
	yScale := float64(srcHeight) / float64(height)

	// Perform bilinear interpolation
	for y := range height {
		for x := range width {
			// Map the target pixel to the source image
			srcX := float64(x) * xScale
			srcY := float64(y) * yScale

			// Get the integer and fractional parts of the source coordinates
			x0 := int(srcX)
			y0 := int(srcY)
			x1 := x0 + 1
			y1 := y0 + 1

			// Clamp the coordinates to the source image bounds
			if x1 >= srcWidth {
				x1 = srcWidth - 1
			}
			if y1 >= srcHeight {
				y1 = srcHeight - 1
			}

			// Get the colors of the four surrounding pixels
			c00 := img.At(srcBounds.Min.X+x0, srcBounds.Min.Y+y0)
			c10 := img.At(srcBounds.Min.X+x1, srcBounds.Min.Y+y0)
			c01 := img.At(srcBounds.Min.X+x0, srcBounds.Min.Y+y1)
			c11 := img.At(srcBounds.Min.X+x1, srcBounds.Min.Y+y1)

			// Calculate the weights for each pixel
			dx := srcX - float64(x0)
			dy := srcY - float64(y0)

			// Interpolate the colors
			r, g, b, a := bilinearInterpolate(c00, c10, c01, c11, dx, dy)

			// Set the color in the new image
			newImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	return newImg
}

func bilinearInterpolate(c00, c10, c01, c11 color.Color, dx, dy float64) (uint8, uint8, uint8, uint8) {
	r00, g00, b00, a00 := c00.RGBA()
	r10, g10, b10, a10 := c10.RGBA()
	r01, g01, b01, a01 := c01.RGBA()
	r11, g11, b11, a11 := c11.RGBA()

	// Interpolate the red channel
	r := uint8((1-dx)*(1-dy)*float64(r00>>8) + dx*(1-dy)*float64(r10>>8) + (1-dx)*dy*float64(r01>>8) + dx*dy*float64(r11>>8))
	// Interpolate the green channel
	g := uint8((1-dx)*(1-dy)*float64(g00>>8) + dx*(1-dy)*float64(g10>>8) + (1-dx)*dy*float64(g01>>8) + dx*dy*float64(g11>>8))
	// Interpolate the blue channel
	b := uint8((1-dx)*(1-dy)*float64(b00>>8) + dx*(1-dy)*float64(b10>>8) + (1-dx)*dy*float64(b01>>8) + dx*dy*float64(b11>>8))
	// Interpolate the alpha channel
	a := uint8((1-dx)*(1-dy)*float64(a00>>8) + dx*(1-dy)*float64(a10>>8) + (1-dx)*dy*float64(a01>>8) + dx*dy*float64(a11>>8))

	return r, g, b, a
}

func CropImage(img image.Image, x, y, width, height int) (image.Image, error) {
	if img.Bounds().Dx() < width+x || img.Bounds().Dy() < height+y {
		return nil, fmt.Errorf("out of bounds crop")
	}

	cropRect := image.Rect(x, y, x+width, y+width)
	croppedImage := image.NewRGBA(cropRect)
	draw.Draw(croppedImage, cropRect, img, cropRect.Min, draw.Src)
	return croppedImage, nil
}
