package sprite_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"asteroid/sprite"
	"asteroid/utils"
)

func TestNewAsteroid(t *testing.T) {
	center := utils.Vector2{X: 100, Y: 100}
	radius := 5
	speed := 2.5
	direction := utils.Vector2{X: 1, Y: 0}

	asteroid := sprite.NewAsteroid(center, radius, speed, direction)
	assert := assert.New(t)

	assert.NotNil(asteroid)
	assert.Equal(center, asteroid.Center)
	assert.Equal(radius, asteroid.Radius)
	assert.Equal(speed, asteroid.Speed)
	assert.Equal(direction, asteroid.Direction)
	assert.Equal(false, asteroid.IsDestoryed())
}

func TestAsteroidUpdate(t *testing.T) {
	center := utils.Vector2{X: 100, Y: 100}
	radius := 5
	speed := 400.0
	direction := utils.Vector2{X: 1, Y: 0}
	asteroid := sprite.NewAsteroid(center, radius, speed, direction)
	dt := 1.0 / 60
	expectedX := center.X + direction.X*speed*dt
	expectedY := center.Y + direction.Y*speed*dt

	asteroid.Update()
	assert := assert.New(t)

	assert.InDelta(expectedX, asteroid.Center.X, 0.0001)
	assert.InDelta(expectedY, asteroid.Center.Y, 0.0001)
	assert.Equal(radius, asteroid.Radius)
	assert.Equal(speed, asteroid.Speed)
	assert.Equal(direction, asteroid.Direction)
}

func TesAasteroidGetHitboxCircule(t *testing.T) {
	center := utils.Vector2{X: 100, Y: 100}
	radius := 5
	asteroid := &sprite.Asteroid{
		Circle: sprite.Circle{
			Center: center,
			Radius: radius,
		},
	}
	pos, rad := asteroid.GetHitboxCircule()

	assert.Equal(t, center, pos)
	assert.Equal(t, radius, rad)
}

func TestAsteroidHitboxCollision(t *testing.T) {
	assert := assert.New(t)

	a := &sprite.Asteroid{
		Circle: sprite.Circle{
			Center: utils.Vector2{X: 100, Y: 100},
			Radius: 20,
		},
	}
	// Test cases
	cases := []struct {
		name     string
		hitbox   sprite.Collidable
		collided bool
	}{
		{
			name: "Collision",
			hitbox: mockCollidable{
				Center: utils.Vector2{X: 110, Y: 110},
				Radius: 15,
			},
			collided: true,
		},
		{
			name: "No Collision",
			hitbox: mockCollidable{
				Center: utils.Vector2{X: 200, Y: 200},
				Radius: 10,
			},
			collided: false,
		},
		{
			name: "Edge Collision",
			hitbox: mockCollidable{
				Center: utils.Vector2{X: 130, Y: 100},
				Radius: 10,
			},
			collided: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collided := a.IsCollided(c.hitbox)
			assert.Equal(c.collided, collided)
		})
	}
}

func TestAsteroidDestory(t *testing.T) {
	assert := assert.New(t)

	a := &sprite.Asteroid{
		Circle: sprite.Circle{
			Center: utils.Vector2{X: 100, Y: 100},
			Radius: 20,
		},
	}
	a.Destory()
	assert.Equal(true, a.IsDestoryed())
}
