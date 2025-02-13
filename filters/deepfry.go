package filters

import (
	"image"
	"math"
)

// NewDeepFryImageFilter provides a filter that applies an image filter chain where the level indicates how much
// each of the other values is applied (values scale at different rates)
// Saturation is 1 + (0.1 * Level)
// Contrast 2 * Level
// PosterizationLevels is max(2, 255/(2**lvl-1))
// There is also a Sharpening filter applied at the end
func NewDeepFryImageFilter(level uint8) ImageFilter {
	return (&deepFrier{int(level)}).Apply
}

type deepFrier struct {
	Level int
}

func (df *deepFrier) Apply(src image.Image) image.Image {
	lf64 := float64(df.Level)
	posterizeLvl := int(max(2.0, 256.0/math.Pow(2, lf64)))
	cfc := ColorFilterChain{
		NewModulateColorFilter(0, 1.0+(0.1*lf64), 1.0),
		NewSigmoidalContrastFilter(2.0 * lf64),
		NewPosterizeColorFilter(posterizeLvl),
	}
	ifc := ImageFilterChain{
		cfc.Apply,
		NewSharpenImageFilter(),
	}

	return ifc.Apply(src)
}
