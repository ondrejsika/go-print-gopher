package go_print_gopher

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/qeesung/image2ascii/convert"
)

func PrintGopher() {
	img, _, err := image.Decode(bytes.NewReader(GOPHER_PNG))
	if err != nil {
		log.Fatalln(err)
	}
	converter := convert.NewImageConverter()
	fmt.Print(converter.Image2ASCIIString(img, &convert.DefaultOptions))
}
