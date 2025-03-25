package sprite_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"asteroid/sprite"
	"asteroid/utils"
)

func TestNewBullet(t *testing.T) {
	center := utils.Vector2{X: 100, Y: 100}
	radius := 5
	speed := 2.5
	direction := utils.Vector2{X: 1, Y: 0}

	bullet := sprite.NewBullet(center, radius, speed, direction)
	assert := assert.New(t)

	assert.NotNil(bullet)
	assert.Equal(center, bullet.Center)
	assert.Equal(radius, bullet.Radius)
	assert.Equal(speed, bullet.Speed)
	assert.Equal(direction, bullet.Direction)
}

func TestBulletUpdate(t *testing.T) {
	center := utils.Vector2{X: 100, Y: 100}
	radius := 5
	speed := 400.0
	direction := utils.Vector2{X: 1, Y: 0}
	bullet := sprite.NewBullet(center, radius, speed, direction)
	dt := 1.0 / 60
	expectedX := center.X + direction.X*speed*dt
	expectedY := center.Y + direction.Y*speed*dt

	bullet.Update()
	assert := assert.New(t)

	assert.InDelta(expectedX, bullet.Center.X, 0.0001)
	assert.InDelta(expectedY, bullet.Center.Y, 0.0001)
	assert.Equal(radius, bullet.Radius)
	assert.Equal(speed, bullet.Speed)
	assert.Equal(direction, bullet.Direction)
}

func TestBulletGetHitboxCircule(t *testing.T) {
	center := utils.Vector2{X: 100, Y: 100}
	radius := 5
	bullet := &sprite.Bullet{
		Circle: sprite.Circle{
			Center: center,
			Radius: radius,
		},
	}
	pos, rad := bullet.GetHitboxCircule()

	assert.Equal(t, center, pos)
	assert.Equal(t, radius, rad)
}

func TestBulletHitboxCollision(t *testing.T) {
	assert := assert.New(t)

	b := &sprite.Bullet{
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
			collided := b.IsCollided(c.hitbox)
			assert.Equal(c.collided, collided)
		})
	}
}
