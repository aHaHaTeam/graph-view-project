package physics

import "math"

type Vec2d struct {
	x, y float64
}

func (v Vec2d) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v Vec2d) add(u Vec2d) Vec2d {
	v.x += u.x
	v.y += u.y
	return v
}

func (v Vec2d) subtract(u Vec2d) Vec2d {
	v.x -= u.x
	v.y -= u.y
	return v
}

func (v Vec2d) multiply(a float64) Vec2d {
	v.x *= a
	v.y *= a
	return v
}

func (v Vec2d) divide(a float64) Vec2d {
	v.x /= a
	v.y /= a
	return v
}
