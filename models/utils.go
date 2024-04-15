package models

import "image/color"

func ColorFromInt(value int) color.RGBA {
	b := value % 1000
	value /= 1000
	g := value % 1000
	value /= 1000
	r := value % 1000
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b)}
}

func ColorToInt(value color.Color) int {
	if value == nil {
		return 0
	}
	r, g, b, _ := value.RGBA()
	return int(r*1000*1000 + g*1000 + b)
}
