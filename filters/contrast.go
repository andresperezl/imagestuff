package filters

import (
	"image/color"
	"math"
)

// NewSigmoidalContrast returns a filter that applies a sigmoidal contrast adjustment to an image.
// The 'contrast' parameter controls the steepness of the sigmoidal curve.
// A higher contrast value results in a stronger effect.
func NewSigmoidalContrastFilter(contrast float64) ColorFilter {
	// logistic function with midpoint 0.5.
	logistic := func(x float64) float64 {
		return 1.0 / (1.0 + math.Exp(-contrast*(x-0.5)))
	}
	// Compute logistic at the endpoints to normalize the output.
	minVal := logistic(0.0)
	maxVal := logistic(1.0)
	scale := 1.0 / (maxVal - minVal)
	filter := &sigmoidalContrast{
		contrast: contrast,
		logistic: logistic,
		minVal:   minVal,
		maxVal:   maxVal,
		scale:    scale,
	}
	return filter.Apply
}

type sigmoidalContrast struct {
	contrast float64
	logistic func(float64) float64
	minVal   float64
	maxVal   float64
	scale    float64
}

func (sc *sigmoidalContrast) Apply(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	// Convert 16-bit channel values to normalized float64 [0,1].
	rf := float64(r) / 65535.0
	gf := float64(g) / 65535.0
	bf := float64(b) / 65535.0

	// Apply the logistic (sigmoidal) transform, then normalize.
	rnew := (sc.logistic(rf) - sc.minVal) * sc.scale
	gnew := (sc.logistic(gf) - sc.minVal) * sc.scale
	bnew := (sc.logistic(bf) - sc.minVal) * sc.scale

	// Clamp values (should already be near [0,1])
	rnew = clamp01f64(rnew)
	gnew = clamp01f64(gnew)
	bnew = clamp01f64(bnew)

	// Convert back to 8-bit and set pixel.
	r8 := uint8(rnew * 255.0)
	g8 := uint8(gnew * 255.0)
	b8 := uint8(bnew * 255.0)
	// Convert 16-bit alpha to 8-bit.
	a8 := uint8(a >> 8)
	return color.NRGBA{R: r8, G: g8, B: b8, A: a8}
}
