package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	_ "image/png" // register PNG decoder
	_ "image/gif" // register GIF decoder
	"io"
	"os"
)

func main() {
	outputKind := flag.String("output", "jpeg", "")
	flag.Parse()
	img, kind, err := image.Decode(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind, "Output format =", *outputKind)
	switch *outputKind {
	case "jpeg":
		err = toJPEG(img, os.Stdout)
	case "png":
		err = toPNG(img, os.Stdout)
	case "gif":
		err = toGIF(img, os.Stdout)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", kind, err)
		os.Exit(1)
	}
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{NumColors: 18})
}
