package filters

import (
	"image/color"
	"testing"
)

func TestModulateApply(t *testing.T) {
	c := color.NRGBA{R: 255, G: 0, B: 0, A: 0xff}
	f := NewModulateColorFilter(0, 1.0, 1.0)
	nc := f(c)
	if nc != c {
		t.Errorf("Expected: %+v,\nGot: %+v", c, nc)
	}
}
