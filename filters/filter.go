package filters

import (
	"cmp"
	"image"
	"image/color"
)

type Filter interface {
	Apply(image.Image) image.Image
}

type ColorFilter func(src color.Color) color.Color

type ImageFilter func(src image.Image) image.Image

type ImageFilterChain []ImageFilter

func (ifc ImageFilterChain) Apply(src image.Image) image.Image {
	result := src
	for _, s := range ifc {
		result = s(result)
	}
	return result
}

type ColorFilterChain []ColorFilter

func (cfc ColorFilterChain) Apply(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := src.At(x, y)
			for _, f := range cfc {
				c = f(c)
			}
			dst.Set(x, y, c)
		}
	}
	return dst
}

func clampFunc[T cmp.Ordered](min, max T) func(T) T {
	return func(v T) T {
		if v < min {
			return min
		}
		if v > max {
			return max
		}
		return v
	}
}

var clamp01f64 = clampFunc(0.0, 1.0)
