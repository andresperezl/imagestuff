package filters

import (
	"image/color"
	"math"

	iscolor "github.com/andresperezl/imagestuff/color"
)

// NewModulateColorFilter create a new filter that  adjusts the brightness, saturation, and hue of the input image.
// The brightness and saturation parameters are multipliers (e.g., 1 means no change, and values < 0 are clamped to 0),
// and the hue parameter is a shift in degrees. A hue value of 0 or 360 produces no change.
func NewModulateColorFilter(hue, saturation, brightness float64) ColorFilter {
	filter := modulate{
		hue:        hue,
		saturation: saturation,
		brightness: brightness,
	}
	return filter.Apply
}

type modulate struct {
	hue        float64
	saturation float64
	brightness float64
}

func (mi *modulate) Apply(c color.Color) color.Color {
	original := color.NRGBAModel.Convert(c).(color.NRGBA)

	// Convert RGB (0-255) to HSL.
	hslColor := iscolor.HSLModel.Convert(c).(iscolor.HSL)
	h, s, l := hslColor.H, hslColor.S, hslColor.L

	// Apply brightness and saturation factors.
	l *= mi.brightness
	l = clamp01f64(l)
	s *= mi.saturation
	s = clamp01f64(s)

	// Apply hue shift.
	h += mi.hue
	// Wrap hue within [0, 360)
	h = math.Mod(h, 360)
	if h < 0 {
		h += 360
	}

	hslColor.H = h
	hslColor.S = s
	hslColor.L = l

	// We need to convert back as HSL doesnt have an Alpha channel
	newColor := color.NRGBAModel.Convert(hslColor).(color.NRGBA)
	newColor.A = original.A
	return newColor
}
