package physics

import "math"

type Vec2d struct {
	x, y float64
}

func vectorFromTwoPoints(begin, end Vec2d) *Vec2d {
	vec := Vec2d{x: end.x - begin.x,
		y: end.y - begin.y}
	return &vec
}

func (vec *Vec2d) length() float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y)
}
