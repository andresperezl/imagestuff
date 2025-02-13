package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"strconv"

	"github.com/andresperezl/imagestuff/filters"
)

func main() {
	lvl, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	path := os.Args[2]
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out, err := os.OpenFile(fmt.Sprintf("test%d.png", lvl), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	filter := filters.NewDeepFryImageFilter(uint8(lvl))
	ifc := filters.ImageFilterChain{
		filters.ColorFilterChain{filters.NewTintColorFilter(color.NRGBA{255, 0, 0, 255}, 0.05)}.Apply,
		filter,
	}
	deepFried := ifc.Apply(src)

	if err := png.Encode(out, deepFried); err != nil {
		panic(err)
	}
}
