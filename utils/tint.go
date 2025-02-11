package utils

import (
	"image"
	"image/color"
)

const (
	max uint32 = 0xffff
)

// Tint applies a color tint to an image by a factor of str/255 (alpha values are kept the same)
func Tint(img image.Image, c color.Color, str uint8) image.Image {
	s := uint32(str)
	s |= s << 8
	t := image.NewRGBA(img.Bounds())
	tr, tg, tb, _ := c.RGBA()
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			nc := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			r, g, b, a := nc.RGBA()
			nr := r*(max-s)/max + tr*s/max
			nc.R = uint8(min(nr, 0xffff) >> 8)
			ng := g*(max-s)/max + tg*s/max
			nc.G = uint8(min(ng, 0xffff) >> 8)
			nb := b*(max-s)/max + tb*s/max
			nc.B = uint8(min(nb, 0xffff) >> 8)
			nc.A = uint8(a >> 8)
			t.Set(x, y, nc)
		}
	}
	return t
}

func min(a, b uint32) uint32 {
	if a > b {
		return b
	}
	return a
}
