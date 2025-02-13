package filters

import (
	"image/color"
	"testing"
)

func TestTintApply(t *testing.T) {
	wc := color.NRGBAModel.Convert(color.White).(color.NRGBA)
	red := color.NRGBA{R: 255, B: 0, G: 0, A: 255}
	f := NewTintColorFilter(red, 0)
	nc := f(wc)
	if wc != nc {
		t.Errorf("Expected: %+v, Got: %+v", wc, nc)
	}
	f = NewTintColorFilter(red, 1.0)
	nc = f(wc)
	if red != nc {
		t.Errorf("Expected: %+v, Got: %+v", red, nc)
	}
}
