package filters

import (
	"image/color"
)

const maxU16 uint32 = 0xffff

// NewTintColorFilter provides a filter that applies a tint color, where factor
// how much of it is applied, a factor of 1 completely replaces the color, a factor
// of 0 does nothing
func NewTintColorFilter(c color.Color, factor float64) ColorFilter {

	s := uint32(factor * 255.0)
	s |= s << 8
	tr, tg, tb, _ := c.RGBA()
	filter := tint{
		rf: tr * s / maxU16,
		gf: tg * s / maxU16,
		bf: tb * s / maxU16,
		of: (maxU16 - s) / maxU16,
	}
	return filter.Apply
}

type tint struct {
	rf, gf, bf uint32
	of         uint32
}

func (t *tint) Apply(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	nc := color.NRGBA{}
	nr := r*t.of + t.rf
	nc.R = uint8(min(nr, 0xffff) >> 8)
	ng := g*t.of + t.gf
	nc.G = uint8(min(ng, 0xffff) >> 8)
	nb := b*t.of + t.bf
	nc.B = uint8(min(nb, 0xffff) >> 8)
	nc.A = uint8(a >> 8)
	return nc
}
