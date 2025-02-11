package utils

import (
	"image"
	"image/color"
	"math"
)

func Tint(img image.Image, c color.Color, str float64) image.Image {
	tinted := image.NewRGBA(img.Bounds())
	tr, tg, tb, ta := c.RGBA()
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Min.Y; y++ {
			nc := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			r, g, b, a := nc.RGBA()
			nr := float64(r)*(1.0-str) + float64(tr)*str
			nc.R = uint8(math.Min(nr, 255))
			ng := float64(g)*(1.0-str) + float64(tg)*str
			nc.G = uint8(math.Min(ng, 255))
			nb := float64(b)*(1.0-str) + float64(tb)*str
			nc.B = uint8(math.Min(nb, 255))
			na := float64(a)*(1.0-str) + float64(ta)*str
			nc.A = uint8(math.Min(na, 255))
			tinted.SetRGBA(x, y, nc)
		}
	}
	return tinted
}
