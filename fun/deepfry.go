package fun

import (
	"image"
	"image/jpeg"
	"io"
	"math"

	"github.com/andresperezl/imagestuff/hsl"
)

type DeepFryOptions struct {
	Hue        float64
	Saturation float64
	Lightness  float64
	Quality    int
}

func DeepFry(out io.Writer, img image.Image, level int, opts *DeepFryOptions) error {
	if level <= 0 || out == nil {
		return nil
	}

	if opts == nil {
		opts = &DeepFryOptions{
			Saturation: 0.1,
			Lightness:  0.1,
			Hue:        36.0,
		}
	}

	newImg := image.NewRGBA(img.Bounds())
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			c := hsl.HSLModel.Convert(img.At(x, y)).(hsl.HSL)
			c.S = math.Min(1.0, c.S+(c.S*opts.Saturation*float64(level)))
			c.L = math.Min(1.0, c.L+(c.L*opts.Lightness*float64(level)))
			c.H = math.Mod(c.H+opts.Hue*float64(level), 360.0)
			if c.H < 0 {
				c.H += 360.0
			}
			newImg.Set(x, y, c)
		}
	}
	q := 100 - opts.Quality*level
	if q < 1 {
		q = 1
	}
	if err := jpeg.Encode(out, newImg, &jpeg.Options{Quality: q}); err != nil {
		return err
	}
	return nil
}
