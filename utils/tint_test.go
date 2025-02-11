package utils

import (
	"image"
	"image/color"
	"testing"
)

func TestTint(t *testing.T) {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 1,
			Y: 1,
		},
	})
	wr, wg, wb, wa := color.White.RGBA()
	img.Set(img.Bounds().Min.X, img.Bounds().Min.Y, color.White)
	timg := Tint(img, color.RGBA{255, 0, 0, 255}, 0)
	tr, tg, tb, ta := timg.At(0, 0).RGBA()
	if wr != tr || wg != tg || wb != tb || wa != ta {
		t.Errorf("Expected: %d,%d,%d,%d, Got: %d,%d,%d,%d", wr, wg, wb, wa, tr, tg, tb, ta)
	}
	red := color.RGBA{R: 255, B: 0, G: 0, A: 255}
	rr, rg, rb, ra := red.RGBA()

	timg = Tint(img, red, 255)
	tr, tg, tb, ta = timg.At(0, 0).RGBA()
	if rr != tr || rg != tg || rb != tb || ra != ta {
		t.Errorf("Expected: %d,%d,%d,%d, Got: %d,%d,%d,%d", rr, rg, rb, ra, tr, tg, tb, ta)
	}
}
