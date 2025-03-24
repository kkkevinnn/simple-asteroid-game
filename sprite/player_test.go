package sprite_test

import (
	"image"
	"testing"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"

	"asteroid/sprite"
	"asteroid/utils"
)

func TestMove(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		name      string
		center    utils.Vector2
		direction sprite.MoveDirection
		distance  float64
		want      utils.Vector2
	}{
		{"move forward", utils.Vector2{X: 0, Y: 0}, sprite.MoveForward, 1, utils.Vector2{X: 0, Y: -1}},
		{"move backward", utils.Vector2{X: 0, Y: 0}, sprite.MoveBackward, 1, utils.Vector2{X: 0, Y: 1}},
	}

	p := &sprite.Player{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p.Center = c.center
			p.Move(c.direction, c.distance)
			assert.Equal(c.want, p.Center)
		})
	}
}

func TestRotate(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		name      string
		rotation  float64
		direction sprite.RotateDirection
		deg       float64
		want      float64
	}{
		{"rotate clockwise", 0, sprite.RotateClockwise, 90, -90},
		{"rotate anticlockwise", 0, sprite.RotateAntiClockwise, 90, 90},
	}

	p := &sprite.Player{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p.Rotation = c.rotation
			p.Rotate(c.direction, c.deg)
			assert.Equal(c.want, p.Rotation)
		})
	}
}

func TestNewPlayer(t *testing.T) {
	assert := assert.New(t)

	center := utils.Vector2{X: 100, Y: 100}
	radius := 20
	bounds := image.Rect(0, 0, 640, 480)
	speed := 5.0
	rotationSpeed := 3.0
	gunConfig := sprite.GunConfig{Radius: 5, Speed: 10.0, RateLimit: time.Millisecond * 500}

	player := sprite.NewPlayer(center, radius, bounds, speed, rotationSpeed, gunConfig)

	assert.Equal(center, player.Center)
	assert.Equal(radius, player.Radius)
	assert.Equal(image.Rect(radius, radius, 640-radius, 480-radius), player.Bounds)
	assert.Equal(speed, player.Speed)
	assert.Equal(rotationSpeed, player.RotationSpeed)
	assert.Equal(gunConfig, player.Gun)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	center := utils.Vector2{X: 100, Y: 100}
	radius := 20
	bounds := image.Rect(0, 0, 640, 480)
	speed := 100.0
	rotationSpeed := 3.0
	gun := sprite.GunConfig{Radius: 5, Speed: 10.0, RateLimit: time.Millisecond * 500}

	p := sprite.NewPlayer(center, radius, bounds, speed, rotationSpeed, gun)

	cases := []struct {
		name     string
		keys     []ebiten.Key
		expected utils.Vector2
	}{
		{"move forward", []ebiten.Key{ebiten.KeyW}, utils.Vector2{X: 100, Y: 98}},
		{"move backward", []ebiten.Key{ebiten.KeyS}, utils.Vector2{X: 100, Y: 100}},
		{"rotate anti-clockwise", []ebiten.Key{ebiten.KeyA}, utils.Vector2{X: 100, Y: 100}},
		{"rotate clockwise", []ebiten.Key{ebiten.KeyD}, utils.Vector2{X: 100, Y: 100}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p.Update(c.keys)
			assert.Equal(c.expected, p.Center)
		})
	}

	t.Run("clamping", func(t *testing.T) {
		p.Center = utils.Vector2{X: float64(radius - 1), Y: float64(radius - 1)}
		p.Update([]ebiten.Key{})
		assert.Equal(utils.Vector2{X: float64(radius), Y: float64(radius)}, p.Center)

		p.Center = utils.Vector2{X: float64(640 - radius + 1), Y: float64(480 - radius + 1)}
		p.Update([]ebiten.Key{})
		assert.Equal(utils.Vector2{X: float64(640 - radius), Y: float64(480 - radius)}, p.Center)
	})
}

func TestTriangle(t *testing.T) {
	assert := assert.New(t)

	center := utils.Vector2{X: 100, Y: 100}
	radius := 30
	bounds := image.Rect(0, 0, 640, 480)
	rotation := float64(0)

	p := &sprite.Player{
		Center:   center,
		Radius:   radius,
		Bounds:   bounds,
		Rotation: rotation,
	}

	cases := []struct {
		name     string
		rotation float64
		vertices [3]*utils.Vector2
	}{
		{
			name:     "Rotation 0",
			rotation: 0,
			vertices: [3]*utils.Vector2{
				{X: 100, Y: 70},
				{X: 120, Y: 130},
				{X: 80, Y: 130},
			},
		},
		{
			name:     "Rotation 90",
			rotation: 90,
			vertices: [3]*utils.Vector2{
				{X: 130, Y: 100},
				{X: 70, Y: 120},
				{X: 70, Y: 80},
			},
		},
		{
			name:     "Rotation 180",
			rotation: 180,
			vertices: [3]*utils.Vector2{
				{X: 100, Y: 130},
				{X: 80, Y: 70},
				{X: 120, Y: 70},
			},
		},
		{
			name:     "Rotation 270",
			rotation: 270,
			vertices: [3]*utils.Vector2{
				{X: 70, Y: 100},
				{X: 130, Y: 80},
				{X: 130, Y: 120},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p.Rotation = c.rotation
			assert.Equal(c.vertices, p.Triangle())
		})
	}

	t.Run("Rotation 45", func(t *testing.T) {
		p.Rotation = 45
		expected := [3]*utils.Vector2{
			{X: 121.21320, Y: 78.786796564},
			{X: 92.9289321, Y: 135.3553390},
			{X: 64.6446609, Y: 107.0710678},
		}
		for i, v := range p.Triangle() {
			assert.InDelta(expected[i].X, v.X, 0.0001)
			assert.InDelta(expected[i].Y, v.Y, 0.0001)
		}
	})
}

func TestGetHitboxCircle(t *testing.T) {
	assert := assert.New(t)

	center := utils.Vector2{X: 100, Y: 100}
	radius := 20
	p := &sprite.Player{
		Center: center,
		Radius: radius,
	}

	pos, rad := p.GetHitboxCircule()

	assert.Equal(center, pos)
	assert.Equal(radius, rad)
}

func TestIsCollided(t *testing.T) {
	assert := assert.New(t)

	p := &sprite.Player{
		Center: utils.Vector2{X: 100, Y: 100},
		Radius: 20,
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
			collided: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			collided := p.IsCollided(c.hitbox)
			assert.Equal(c.collided, collided)
		})
	}
}

func TestFire(t *testing.T) {
	assert := assert.New(t)
	gunConfig := sprite.GunConfig{Radius: 5, Speed: 10.0, RateLimit: time.Millisecond * 500}
	p := &sprite.Player{
		Center: utils.Vector2{X: 100, Y: 100},
		Radius: 20,
		Gun:    gunConfig,
	}

	// First fire
	bullet, err := p.Fire()
	assert.NoError(err)
	assert.NotNil(bullet)

	// Try to fire again before rate limit
	_, err = p.Fire()
	assert.Equal(sprite.ErrGunNotReady, err)

	// Wait for rate limit
	time.Sleep(gunConfig.RateLimit)

	// Fire again after rate limit
	bullet, err = p.Fire()
	assert.NoError(err)
	assert.NotNil(bullet)
}
