package filters

import (
	"image/color"
	"math"
)

func NewPosterizeColorFilter(levels int) ColorFilter {
	// Calculate the step size between quantized values.
	step := 255.0 / float64(levels-1)

	// quantize maps a channel value (0-255) to its posterized equivalent.
	quantize := func(val float64) uint8 {
		quantized := math.Round(val/step) * step
		if quantized < 0 {
			quantized = 0
		}
		if quantized > 255 {
			quantized = 255
		}
		return uint8(quantized)
	}
	filter := posterize{
		levels:   levels,
		step:     step,
		quantize: quantize,
	}
	return filter.Apply
}

type posterize struct {
	levels   int
	step     float64
	quantize func(float64) uint8
}

func (p *posterize) Apply(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	// Convert 16-bit values to 8-bit.
	r8 := float64(uint8(r >> 8))
	g8 := float64(uint8(g >> 8))
	b8 := float64(uint8(b >> 8))
	a8 := uint8(a >> 8)

	rnew := p.quantize(r8)
	gnew := p.quantize(g8)
	bnew := p.quantize(b8)

	return color.NRGBA{R: rnew, G: gnew, B: bnew, A: a8}
}
