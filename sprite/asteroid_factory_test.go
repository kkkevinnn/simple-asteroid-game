package sprite_test

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"

	"asteroid/sprite"
)

func TestAsteroidNewFactory(t *testing.T) {
	minRadius := 20
	kind := 3
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	maxSpeed := 100.0
	minSpeed := 40.0
	maxAngle := 30.0

	factory := sprite.NewAsteroidFactory(minRadius, kind, bounds, maxSpeed, minSpeed, maxAngle)

	assert := assert.New(t)
	assert.NotNil(factory)
	assert.Equal(minRadius, factory.MinRadius)
	assert.Equal(kind, factory.Kind)
	assert.Equal(bounds, factory.Bounds)
	assert.Equal(maxSpeed, factory.MaxSpeed)
	assert.Equal(minSpeed, factory.MinSpeed)
	assert.Equal(maxAngle, factory.MaxAngle)
}

func TestAsteroidNewAestroid(t *testing.T) {
	factory := sprite.NewAsteroidFactory(20, 3, image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}, 100.0, 40.0, 30.0)

	type expected struct {
		radius  [3]int
		speed   [2]float64
		centerX [2]float64
		centerY [2]float64
	}

	cases := []struct {
		name     string
		edge     int
		expected expected
	}{
		{
			name: "North",
			edge: 0,
			expected: expected{
				radius:  [3]int{20, 40, 60},
				speed:   [2]float64{40.0, 100.0},
				centerX: [2]float64{0.0, 1000.0},
				centerY: [2]float64{0.0, 0.0},
			},
		},
		{
			name: "East",
			edge: 1,
			expected: expected{
				radius:  [3]int{20, 40, 60},
				speed:   [2]float64{40.0, 100.0},
				centerX: [2]float64{1000.0, 1000.0},
				centerY: [2]float64{0.0, 1000.0},
			},
		},
		{
			name: "South",
			edge: 2,
			expected: expected{
				radius:  [3]int{20, 40, 60},
				speed:   [2]float64{40.0, 100.0},
				centerX: [2]float64{0.0, 1000.0},
				centerY: [2]float64{1000.0, 1000.0},
			},
		},
		{
			name: "West",
			edge: 3,
			expected: expected{
				radius:  [3]int{20, 40, 60},
				speed:   [2]float64{40.0, 100.0},
				centerX: [2]float64{0.0, 0.0},
				centerY: [2]float64{0.0, 1000.0},
			},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			asteroid := factory.NewAsteroid(c.edge)
			assert.NotNil(asteroid)
			assert.Contains(c.expected.radius, asteroid.Radius)

			assert.LessOrEqual(c.expected.speed[0], asteroid.Speed, "edge %d: speed min", c.edge)
			assert.LessOrEqual(c.expected.centerX[0]-float64(asteroid.Radius), asteroid.Center.X, "edge %d: center X min", c.edge)
			assert.LessOrEqual(c.expected.centerY[0]-float64(asteroid.Radius), asteroid.Center.Y, "edge %d: center Y min", c.edge)

			assert.GreaterOrEqual(c.expected.speed[1], asteroid.Speed, "edge %d: speed max", c.edge)
			assert.GreaterOrEqual(c.expected.centerX[1]+float64(asteroid.Radius), asteroid.Center.X, "edge %d: center X max", c.edge)
			assert.GreaterOrEqual(c.expected.centerY[1]+float64(asteroid.Radius), asteroid.Center.Y, "edge %d: center Y max", c.edge)
		})
	}
}
