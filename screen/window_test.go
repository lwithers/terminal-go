package screen

import "testing"

func BenchmarkDrawRune(b *testing.B) {
	win := NewWindow(100, 100)
	w, h := win.Size()
	for i := 0; i < b.N; i++ {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				win.DrawRune('x', White, Black, x, y)
			}
		}
	}
}

func BenchmarkDrawRunes(b *testing.B) {
	win := NewWindow(100, 100)
	w, h := win.Size()
	r := make([]rune, w)
	for i := range r {
		r[i] = '·'
	}
	for i := 0; i < b.N; i++ {
		for y := 0; y < h; y++ {
			win.DrawRunes(r, White, Black, 0, y)
		}
	}
}
