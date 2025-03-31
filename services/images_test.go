package services

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func newSampleImage() image.Image {
	baseImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	draw.Draw(baseImg, baseImg.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.Point{}, draw.Src)
	return baseImg
}

func TestResizeImage(t *testing.T) {
	baseImage := newSampleImage()

	// resize to lower size
	resizedImage := ResizeImage(baseImage, 5, 5)
	if resizedImage.Bounds().Dx() != 5 || resizedImage.Bounds().Dy() != 5 {
		t.Errorf("expected resized image dimensions to be 5x5, got %dx%d", resizedImage.Bounds().Dx(), resizedImage.Bounds().Dy())
	}

	// resize to larger size
	resizedImage = ResizeImage(baseImage, 20, 20)
	if resizedImage.Bounds().Dx() != 20 || resizedImage.Bounds().Dy() != 20 {
		t.Errorf("expected resized image dimensions to be 20x20, got %dx%d", resizedImage.Bounds().Dx(), resizedImage.Bounds().Dy())
	}

	// resize to the same size
	resizedImage = ResizeImage(baseImage, 10, 10)
	if resizedImage.Bounds().Dx() != 10 || resizedImage.Bounds().Dy() != 10 {
		t.Errorf("expected resized image dimensions to be 10x10, got %dx%d", resizedImage.Bounds().Dx(), resizedImage.Bounds().Dy())
	}
}

func TestCropImage(t *testing.T) {
	baseImage := newSampleImage()

	// test valid crop
	croppedImage, err := CropImage(baseImage, 2, 2, 5, 5)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if croppedImage.Bounds().Dx() != 5 || croppedImage.Bounds().Dy() != 5 {
		t.Errorf("expected cropped image dimensions to be 5x5, got %v:%v", croppedImage.Bounds().Dx(), croppedImage.Bounds().Dy())
	}

	// crop out of bonds error
	_, err = CropImage(baseImage, 8, 8, 5, 5)
	if err == nil {
		t.Errorf("expected an error for out of bonds crop, but got none")
	}

	// crop entire image
	croppedImage, err = CropImage(baseImage, 0, 0, 10, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if croppedImage.Bounds().Dx() != 10 || croppedImage.Bounds().Dy() != 10 {
		t.Errorf("expected cropped image dimensions to be 10x10, got %dx%d", croppedImage.Bounds().Dx(), croppedImage.Bounds().Dy())
	}
}
