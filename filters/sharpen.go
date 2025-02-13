package filters

import (
	"image"
	"image/color"
)

var kernel = [3][3]float64{
	{0, -1, 0},
	{-1, 5, -1},
	{0, -1, 0},
}

// NewSharpenImageFilter creates a filter that applies a simple convolution-based sharpening filter using a 3x3 kernel.
// The kernel used here is a common sharpening filter.
func NewSharpenImageFilter() ImageFilter {
	return sharpen
}

var clamp255f64 = clampFunc(0.0, 255.0)

func sharpen(src image.Image) image.Image {
	dst := image.NewRGBA(src.Bounds())
	bounds := src.Bounds()
	clampX := clampFunc(bounds.Min.X, bounds.Max.X-1)
	clampY := clampFunc(bounds.Min.Y, bounds.Max.Y-1)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var rSum, gSum, bSum float64
			// Apply the kernel to each neighboring pixel.
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					nx := clampX(x + kx)
					ny := clampY(y + ky)
					r, g, b, _ := src.At(nx, ny).RGBA()
					r8 := float64(uint8(r >> 8))
					g8 := float64(uint8(g >> 8))
					b8 := float64(uint8(b >> 8))
					weight := kernel[ky+1][kx+1]
					rSum += r8 * weight
					gSum += g8 * weight
					bSum += b8 * weight
				}
			}
			rNew := uint8(clamp255f64(rSum))
			gNew := uint8(clamp255f64(gSum))
			bNew := uint8(clamp255f64(bSum))
			_, _, _, a := src.At(x, y).RGBA()
			a8 := uint8(a >> 8)
			dst.Set(x, y, color.NRGBA{R: rNew, G: gNew, B: bNew, A: a8})
		}
	}
	return dst
}
