package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
)

func expandRect(a, b sdl.Rect) sdl.Rect {
	r := sdl.Rect{
		X: min(a.X, b.X),
		Y: min(a.Y, b.Y),
	}
	r.W = max(a.X+a.W, b.X+b.W) - r.X
	r.H = max(a.Y+a.H, b.Y+b.H) - r.Y
	return r
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func growRect(rect sdl.Rect, top, right, bottom, left int32) sdl.Rect {
	return sdl.Rect{
		X: rect.X - left,
		Y: rect.Y - top,
		W: rect.W + left + right,
		H: rect.H + top + bottom,
	}
}
