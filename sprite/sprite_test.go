package sprite_test

import (
	"asteroid/sprite"
	"image"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClampPoint(t *testing.T) {
	assert := assert.New(t)

	testBox := image.Rect(0, 0, 10, 10)
	cases := []struct {
		name     string
		point    sprite.Point
		expected sprite.Point
	}{
		{
			name:     "Point inside box",
			point:    sprite.Point{X: 5, Y: 5},
			expected: sprite.Point{X: 5, Y: 5},
		},
		{
			name:     "Point over box",
			point:    sprite.Point{X: 5, Y: -5},
			expected: sprite.Point{X: 5, Y: 0},
		},
		{
			name:     "Point under box",
			point:    sprite.Point{X: 5, Y: 15},
			expected: sprite.Point{X: 5, Y: 10},
		},
		{
			name:     "Point left of box",
			point:    sprite.Point{X: -5, Y: 5},
			expected: sprite.Point{X: 0, Y: 5},
		},
		{
			name:     "Point right of box",
			point:    sprite.Point{X: 15, Y: 5},
			expected: sprite.Point{X: 10, Y: 5},
		},
		{
			name:     "Point top left of box",
			point:    sprite.Point{X: -5, Y: -5},
			expected: sprite.Point{X: 0, Y: 0},
		},
		{
			name:     "Point top right of box",
			point:    sprite.Point{X: 15, Y: -5},
			expected: sprite.Point{X: 10, Y: 0},
		},
		{
			name:     "Point bottom left of box",
			point:    sprite.Point{X: -5, Y: 15},
			expected: sprite.Point{X: 0, Y: 10},
		},
		{
			name:     "Point bottom right of box",
			point:    sprite.Point{X: 15, Y: 15},
			expected: sprite.Point{X: 10, Y: 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.point.Clamp(testBox)
			assert.Equal(c.expected, c.point, "The point should be clamped to the box")
		})
	}
}
