package hsl

import (
	"image/color"
	"math"
)

type HSL struct {
	H, S, L float64
}

func (c HSL) RGBA() (r, g, b, a uint32) {
	return c.f(0), c.f(8), c.f(4), 0xffff
}

func (c HSL) f(n int) uint32 {
	k := c.k(n)
	mink := math.Min(k-3.0, math.Min(9.0-k, 1))
	amax := c.a() * math.Max(-1, mink)
	v := uint32(math.Round(255.0 * (c.L - amax)))
	return (v << 8) | v
}

func (c HSL) k(n int) float64 {
	return math.Mod(float64(n)+float64(c.H)/30.0, 12.0)
}

func (c HSL) a() float64 {
	return c.S * math.Min(c.L, 1.0-c.L)
}

var HSLModel color.Model = color.ModelFunc(hslModel)

func hslModel(c color.Color) color.Color {

	if _, ok := c.(HSL); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	rprime := float64(r>>8) / 255.0
	gprime := float64(g>>8) / 255.0
	bprime := float64(b>>8) / 255.0

	xmax := math.Max(rprime, math.Max(gprime, bprime))
	xmin := math.Min(rprime, math.Min(gprime, bprime))

	chroma := xmax - xmin
	l := (xmax + xmin) / 2.0
	var h float64
	if chroma == 0 {
		h = 0
	} else {
		switch xmax {
		case rprime:
			h = 60.0 * (gprime - bprime) / chroma
		case gprime:
			h = 60.0 * (2.0 + (bprime-rprime)/chroma)
		case bprime:
			h = 60.0 * (4.0 + (rprime-gprime)/chroma)
		}
	}
	if h < 0.0 {
		h += 360.0
	}

	var s float64
	if l == 0.0 || l == 1.0 {
		s = 0
	} else {
		s = (xmax - l) / math.Min(l, 1.0-l)
	}

	return HSL{h, s, l}
}
