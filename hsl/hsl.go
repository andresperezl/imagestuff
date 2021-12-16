package hsl

import (
	"image/color"
	"math"
)

const (
	oneThird = 1.0 / 3.0
	twoThird = 2.0 / 3.0
	twoPi    = 2 * math.Pi
)

type HSL struct {
	H, S, L float64
}

func (c HSL) RGBA() (r, g, b, a uint32) {
	if c.S == 0 {
		gray := uint32(math.Round(255.0 * c.L))
		return gray, gray, gray, 0xffff
	}

	var t1, t2 float64
	if c.L < 0.5 {
		t1 = c.L * (1.0 + c.S)
	} else {
		t1 = c.L + c.S - c.L*c.S
	}

	t2 = 2*c.L - t1

	hp := math.Remainder(c.H, twoPi) / twoPi
	tR := hp + oneThird
	tG := hp
	tB := hp - oneThird
	if tR > 1 {
		tR -= 1
	}
	if tB < 0 {
		tB += 1
	}

	r = getChannel(tR, t1, t2)
	g = getChannel(tG, t1, t2)
	b = getChannel(tB, t1, t2)

	return r, g, b, 0xffff
}

func getChannel(tc, t1, t2 float64) uint32 {
	v := tc
	if 6.0*tc < 1.0 {
		v = t2 + (t1-t2)*6.0*tc
	} else if 2.0*tc < 1.0 {
		v = t1
	} else if 3.0*tc < 2.0 {
		v = t2 + (t1-t2)*(twoThird-tc)*6.0
	}
	return uint32(math.Round(v * 255.0))
}

var HSLModel color.Model = color.ModelFunc(hslModel)

func hslModel(c color.Color) color.Color {

	if _, ok := c.(HSL); ok {
		return c
	}
	var h, s, l float64
	r, g, b, _ := c.RGBA()
	rp := float64(r) / 255.0
	gp := float64(g) / 255.0
	bp := float64(b) / 255.0
	min := math.Min(rp, math.Min(gp, bp))
	max := math.Max(rp, math.Max(gp, bp))
	l = (min + max) / 2
	if min == max {
		return HSL{0, 0, l}
	}
	mmm := max - min
	if l <= 0.5 {
		s = mmm / (max + min)
	} else {
		s = mmm / (2.0 - max - min)
	}
	switch max {
	case rp:
		h = (gp - bp) / mmm
	case gp:
		h = 2.0 + (bp-rp)/mmm
	case bp:
		h = 4.0 + (rp-gp)/mmm
	}
	if h < 0 {
		h += twoPi
	}
	return HSL{h, s, l}
}
