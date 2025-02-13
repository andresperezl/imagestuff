package color

import (
	"image/color"
	"math"
)

type HSL struct {
	H, S, L float64
}

func (c HSL) RGBA() (uint32, uint32, uint32, uint32) {
	var r, g, b float64

	if c.S == 0 {
		// Achromatic color (gray).
		r, g, b = c.L, c.L, c.L
	} else {
		var q float64
		if c.L < 0.5 {
			q = c.L * (1 + c.S)
		} else {
			q = c.L + c.S - c.L*c.S
		}
		p := 2*c.L - q
		hk := c.H / 360

		// Helper function to convert hue to rgb.

		r = hue2rgb(p, q, hk+1.0/3.0)
		g = hue2rgb(p, q, hk)
		b = hue2rgb(p, q, hk-1.0/3.0)
	}

	return uint32(math.Round(r * 0xffff)), uint32(math.Round(g * 0xffff)), uint32(math.Round(b * 0xffff)), 0xffff
}

func hue2rgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

var HSLModel color.Model = color.ModelFunc(hslModel)

func hslModel(c color.Color) color.Color {

	if _, ok := c.(HSL); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	rf := float64(r) / 65535.0
	gf := float64(g) / 65535.0
	bf := float64(b) / 65535.0

	max := max(rf, max(gf, bf))
	min := min(rf, min(gf, bf))
	l := (max + min) / 2

	var h, s float64
	if max == min {
		h, s = 0, 0 // achromatic
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case rf:
			h = (gf - bf) / d
			if gf < bf {
				h += 6
			}
		case gf:
			h = (bf-rf)/d + 2
		case bf:
			h = (rf-gf)/d + 4
		}
		h *= 60
	}

	return HSL{h, s, l}
}
