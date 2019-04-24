package main

import (
	"image/jpeg"
	"os"

	"github.com/chai2010/webp"
)

func main() {

	in, err := os.Open("./in/1.jpg")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	img, err := jpeg.Decode(in)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("./out/1.webp")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = webp.Encode(out, img, &webp.Options{Quality: 70})
	if err != nil {
		panic(err)
	}
}
