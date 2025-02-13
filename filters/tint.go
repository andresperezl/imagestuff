package filters

import (
	"image/color"
)

const maxU16 uint32 = 0xffff

// NewTintColorFilter provides a filter that applies a tint color, where factor
// how much of it is applied, a factor of 1 completely replaces the color, a factor
// of 0 does nothing
func NewTintColorFilter(c color.Color, factor float64) ColorFilter {
	tr, tg, tb, _ := c.RGBA()
	f := clamp01f64(factor)
	filter := tint{
		rf: float64(tr) * f,
		gf: float64(tg) * f,
		bf: float64(tb) * f,
		fi: 1.0 - f,
	}
	return filter.Apply
}

type tint struct {
	rf, gf, bf float64
	fi         float64
}

func (t *tint) Apply(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	nc := color.NRGBA{}
	nr := float64(r)*t.fi + t.rf
	nc.R = uint8(uint32(min(nr, 65535.0)) >> 8)
	ng := float64(g)*t.fi + t.gf
	nc.G = uint8(uint32(min(ng, 65535.0)) >> 8)
	nb := float64(b)*t.fi + t.bf
	nc.B = uint8(uint32(min(nb, 65535.0)) >> 8)
	nc.A = uint8(a >> 8)
	return nc
}
