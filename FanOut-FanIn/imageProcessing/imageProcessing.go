package imageProcessing

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(filePath string) image.Image {
	inputFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("error is the path:", err)
		panic(err)
	}

	return img
}

func WriteImage(imagePath string, img image.Image) {
	outPutFile, err := os.Create(imagePath)
	if err != nil {
		panic(err)
	}
	defer outPutFile.Close()

	err = jpeg.Encode(outPutFile, img, nil)
	if err != nil {
		panic(err)
	}
}

func ResizeImage(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}
