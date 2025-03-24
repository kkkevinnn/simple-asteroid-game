package utils

import "math"

type Vector2 struct {
	X, Y float64
}

func NewVector2(x, y float64) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}

// Add adds v2 to v.
func (v *Vector2) Add(v2 Vector2) *Vector2 {
	v.X += v2.X
	v.Y += v2.Y

	return v
}

// Sub subtracts v2 from v.
func (v *Vector2) Sub(v2 Vector2) *Vector2 {
	v.X -= v2.X
	v.Y -= v2.Y

	return v
}

// Rotate rotates the vector by deg degrees. Positive deg means counter-clockwise rotation.
func (v *Vector2) Rotate(deg float64) *Vector2 {
	rad := deg * math.Pi / 180
	x := v.X*math.Cos(rad) - v.Y*math.Sin(rad)
	y := v.X*math.Sin(rad) + v.Y*math.Cos(rad)
	v.X, v.Y = x, y

	return v
}

// Scale scales the vector by s.
func (v *Vector2) Scale(s float64) *Vector2 {
	v.X *= s
	v.Y *= s

	return v
}

// Length returns the length of the vector.
func (v *Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Copy returns a copy of the vector.
func (v *Vector2) Copy() *Vector2 {
	return NewVector2(v.X, v.Y)
}
