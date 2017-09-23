package lib

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const opaque uint16 = 65535

func CommandEncode(fileIn, fileOut, message string) {
	// Encode message into base64
	base64Data := []byte(base64.StdEncoding.EncodeToString([]byte(message)))

	// Open file
	imageReader, err := os.Open(fileIn)
	if err != nil {
		fmt.Println("Could not open file: ", err)
		os.Exit(1)
	}
	defer imageReader.Close()

	// Decode image
	imageDecoder, _, err := image.Decode(imageReader)
	if err != nil {
		fmt.Println("Could not decode image: ", err)
		os.Exit(1)
	}

	bounds := imageDecoder.Bounds()
	imageOut := image.NewRGBA64(bounds)
	i := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var a16 uint16
			r, g, b, _ := imageDecoder.At(x, y).RGBA()

			// Set the alpha channel
			if len(base64Data) > i {
				a16 = opaque - uint16(base64Data[i]) + uint16(1)
			} else {
				a16 = opaque
			}

			c := color.RGBA64{uint16(r), uint16(g), uint16(b), a16}
			imageOut.Set(x, y, c)
			i++
		}
	}

	writer, _ := os.Create(fileOut)

	png.Encode(writer, imageOut)
}

func CommandDecode(fileIn string) {
	// Open file
	imageReader, err := os.Open(fileIn)
	if err != nil {
		fmt.Println("Could not open file: ", err)
		os.Exit(1)
	}
	defer imageReader.Close()

	// Decode image
	imageDecoder, _, err := image.Decode(imageReader)
	if err != nil {
		fmt.Println("Could not decode image: ", err)
		os.Exit(1)
	}

	bounds := imageDecoder.Bounds()
	base64Data := make([]byte, 0)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := imageDecoder.At(x, y).RGBA()
			if uint16(a) != opaque {
				base64Data = append(base64Data, byte(opaque-uint16(a)+uint16(1)))
			}
		}
	}

	message, err := base64.StdEncoding.DecodeString(string(base64Data))
	if err != nil {
		fmt.Println("Could not decode base64: ", err)
		os.Exit(1)
	}

	fmt.Println(string(message))
}
