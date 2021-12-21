package fun

import (
	"fmt"
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

	if out == nil {
		return fmt.Errorf("out cannot be nil")
	}

	if level <= 0 {
		return fmt.Errorf("level needs to be greater than 0")
	}

	if opts == nil {
		opts = &DeepFryOptions{
			Saturation: 0.1,
			Lightness:  0.1,
			Hue:        36,
			Quality:    10,
		}
	}
	lf := float64(level)
	newImg := image.NewRGBA(img.Bounds())
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			c := hsl.HSLModel.Convert(img.At(x-x%level, y-y%level)).(hsl.HSL)
			c.S = math.Min(1.0, c.S+(c.S*opts.Saturation*lf))
			c.L = math.Min(1.0, c.L+(c.L*opts.Lightness*lf))
			c.H = math.Mod(c.H+opts.Hue*lf, 360.0)
			if c.H < 0 {
				c.H += 360.0
			}
			newImg.Set(x, y, c)
		}
	}
	q := math.Max(1.0, 100.0*math.Pow(1.0/float64(opts.Quality), lf))
	if err := jpeg.Encode(out, newImg, &jpeg.Options{Quality: int(q)}); err != nil {
		return err
	}
	return nil
}
