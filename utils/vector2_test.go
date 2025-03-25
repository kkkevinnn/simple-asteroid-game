package utils_test

import (
	"image"
	"math"
	"testing"

	. "asteroid/utils"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	v1 := Vector2{1, 2}
	v2 := Vector2{3, 4}
	v1.Add(v2)
	want := Vector2{4, 6}
	assert.Equal(t, want, v1, "The sum of v1(%v) and v2(%v) should be %v", v1, v2, want)
}

func TestSub(t *testing.T) {
	v1 := Vector2{1, 2}
	v2 := Vector2{3, 4}
	v1.Sub(v2)
	want := Vector2{-2, -2}
	assert.Equal(t, want, v1, "The sub of v1(%v) and v2(%v) should be %v", v1, v2, want)
}

func TestRotate(t *testing.T) {
	cases := []struct {
		name string
		v    Vector2
		deg  float64
		want Vector2
	}{
		{"0 deg", Vector2{1, 0}, 0, Vector2{1, 0}},
		{"90 deg", Vector2{1, 0}, 90, Vector2{0, 1}},
		{"180 deg", Vector2{1, 0}, 180, Vector2{-1, 0}},
		{"270 deg", Vector2{1, 0}, 270, Vector2{0, -1}},
		{"360 deg", Vector2{1, 0}, 360, Vector2{1, 0}},
		{"-90 deg", Vector2{1, 0}, -90, Vector2{0, -1}},
		{"-180 deg", Vector2{1, 0}, -180, Vector2{-1, 0}},
		{"-270 deg", Vector2{1, 0}, -270, Vector2{0, 1}},
		{"-360 deg", Vector2{1, 0}, -360, Vector2{1, 0}},
		{"45 deg", Vector2{1, 0}, 45, Vector2{0.7071067811865476, 0.7071067811865476}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			orig := c.v
			c.v.Rotate(c.deg)
			assert.InDelta(t, c.want.X, c.v.X, 0.000000001, "The rotate of v(%v) by %v should be %v but got %v", orig, c.deg, c.want, c.v)
			assert.InDelta(t, c.want.Y, c.v.Y, 0.000000001, "The rotate of v(%v) by %v should be %v but got %v", orig, c.deg, c.want, c.v)
		})
	}
}

func TestScale(t *testing.T) {
	cases := []struct {
		name string
		v    Vector2
		s    float64
		want Vector2
	}{
		{"1.0", Vector2{1, 2}, 1.0, Vector2{1, 2}},
		{"2.0", Vector2{1, 2}, 2.0, Vector2{2, 4}},
		{"0.5", Vector2{1, 2}, 0.5, Vector2{0.5, 1}},
		{"-1.0", Vector2{1, 2}, -1.0, Vector2{-1, -2}},
		{"-2.0", Vector2{1, 2}, -2.0, Vector2{-2, -4}},
		{"-0.5", Vector2{1, 2}, -0.5, Vector2{-0.5, -1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			orig := c.v
			c.v.Scale(c.s)
			assert.Equal(t, c.want, c.v, "The scale of v(%v) by %v should be %v but got %v", orig, c.s, c.want, c.v)
		})
	}
}

func TestLength(t *testing.T) {
	cases := []struct {
		name string
		v    Vector2
		want float64
	}{
		{"(0, 0)", Vector2{0, 0}, 0},
		{"(1, 0)", Vector2{1, 0}, 1},
		{"(0, 1)", Vector2{0, 1}, 1},
		{"(1, 1)", Vector2{1, 1}, math.Sqrt(2)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, c.v.Length(), "The length of v(%v) should be %v but got %v", c.v, c.want, c.v.Length())
		})
	}
}

func TestCopy(t *testing.T) {
	cases := []struct {
		name string
		v    Vector2
		want Vector2
	}{
		{"(0, 0)", Vector2{0, 0}, Vector2{0, 0}},
		{"(1, 0)", Vector2{1, 0}, Vector2{1, 0}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, *c.v.Copy(), "The copy of v(%v) should be %v but got %v", c.v, c.want, c.v.Copy())
		})
	}
}

func TestClampVector2(t *testing.T) {
	assert := assert.New(t)

	testBox := image.Rect(0, 0, 10, 10)
	cases := []struct {
		name     string
		vector2  Vector2
		expected Vector2
	}{
		{
			name:     "vector2 inside box",
			vector2:  Vector2{X: 5, Y: 5},
			expected: Vector2{X: 5, Y: 5},
		},
		{
			name:     "vector2 over box",
			vector2:  Vector2{X: 5, Y: -5},
			expected: Vector2{X: 5, Y: 0},
		},
		{
			name:     "vector2 under box",
			vector2:  Vector2{X: 5, Y: 15},
			expected: Vector2{X: 5, Y: 10},
		},
		{
			name:     "vector2 left of box",
			vector2:  Vector2{X: -5, Y: 5},
			expected: Vector2{X: 0, Y: 5},
		},
		{
			name:     "vector2 right of box",
			vector2:  Vector2{X: 15, Y: 5},
			expected: Vector2{X: 10, Y: 5},
		},
		{
			name:     "vector2 top left of box",
			vector2:  Vector2{X: -5, Y: -5},
			expected: Vector2{X: 0, Y: 0},
		},
		{
			name:     "vector2 top right of box",
			vector2:  Vector2{X: 15, Y: -5},
			expected: Vector2{X: 10, Y: 0},
		},
		{
			name:     "vector2 bottom left of box",
			vector2:  Vector2{X: -5, Y: 15},
			expected: Vector2{X: 0, Y: 10},
		},
		{
			name:     "vector2 bottom right of box",
			vector2:  Vector2{X: 15, Y: 15},
			expected: Vector2{X: 10, Y: 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.vector2.Clamp(testBox)
			assert.Equal(c.expected, c.vector2, "The point should be clamped to the box")
		})
	}
}
