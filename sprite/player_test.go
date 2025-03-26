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

func TestPlayerMove(t *testing.T) {
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

	p := &sprite.Player{
		Circle: sprite.Circle{
			Direction: utils.Vector2{X: 0, Y: -1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p.Center = c.center
			p.Move(c.direction, c.distance)
			assert.Equal(c.want, p.Center)
		})
	}
}

func TestPlayerRotate(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		name         string
		rotDirection sprite.RotateDirection
		deg          float64
		oriDirection utils.Vector2
		want         utils.Vector2
	}{
		{"rotate clockwise", sprite.RotateClockwise, 90, utils.Vector2{X: 0, Y: -1}, utils.Vector2{X: 1, Y: 0}},
		{"rotate anticlockwise", sprite.RotateAntiClockwise, 90, utils.Vector2{X: 0, Y: -1}, utils.Vector2{X: -1, Y: 0}},
	}

	p := &sprite.Player{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p.Direction = c.oriDirection
			p.Rotate(c.rotDirection, c.deg)
			assert.InDelta(c.want.X, p.Direction.X, 0.00001)
			assert.InDelta(c.want.Y, p.Direction.Y, 0.00001)
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
	assert.Equal(speed, player.Speed)
	assert.Equal(utils.Vector2{X: 0, Y: -1}, player.Direction)
	assert.Equal(image.Rect(radius, radius, 640-radius, 480-radius), player.Bounds)
	assert.Equal(rotationSpeed, player.RotationSpeed)
	assert.Equal(gunConfig, player.Gun)
}

func TestPlayerUpdate(t *testing.T) {
	assert := assert.New(t)

	center := utils.Vector2{X: 100, Y: 100}
	radius := 20
	bounds := image.Rect(0, 0, 640, 480)
	speed := 100.0
	rotationSpeed := 300.0
	gun := sprite.GunConfig{Radius: 5, Speed: 10.0, RateLimit: time.Millisecond * 500}

	p := sprite.NewPlayer(center, radius, bounds, speed, rotationSpeed, gun)

	type Case struct {
		name     string
		keys     []ebiten.Key
		expected utils.Vector2
	}
	moveCases := []Case{
		{"move forward", []ebiten.Key{ebiten.KeyW}, utils.Vector2{X: 100, Y: 98.33333}},
		{"move backward", []ebiten.Key{ebiten.KeyS}, utils.Vector2{X: 100, Y: 100}},
	}

	for _, c := range moveCases {
		t.Run(c.name, func(t *testing.T) {
			p.Update(c.keys)
			assert.InDelta(c.expected.X, p.Center.X, 0.0001)
			assert.InDelta(c.expected.Y, p.Center.Y, 0.0001)
		})
	}

	rotateCases := []Case{
		{"rotate anti-clockwise", []ebiten.Key{ebiten.KeyA}, utils.Vector2{X: -0.08715, Y: -0.99619}},
		{"rotate clockwise", []ebiten.Key{ebiten.KeyD}, utils.Vector2{X: 0, Y: -1}},
	}

	for _, c := range rotateCases {
		t.Run(c.name, func(t *testing.T) {
			p.Update(c.keys)
			assert.InDelta(c.expected.X, p.Direction.X, 0.0001)
			assert.InDelta(c.expected.Y, p.Direction.Y, 0.0001)
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

func TestPlayerTriangle(t *testing.T) {
	assert := assert.New(t)

	center := utils.Vector2{X: 100, Y: 100}
	radius := 30
	bounds := image.Rect(0, 0, 640, 480)

	p := &sprite.Player{
		Circle: sprite.Circle{
			Center: center,
			Radius: radius,
		},
		Bounds: bounds,
	}

	cases := []struct {
		name      string
		direction utils.Vector2
		vertices  [3]*utils.Vector2
	}{
		{
			name:      "Rotation 0",
			direction: utils.Vector2{X: 0, Y: -1},
			vertices: [3]*utils.Vector2{
				{X: 100, Y: 70},
				{X: 120, Y: 130},
				{X: 80, Y: 130},
			},
		},
		{
			name:      "Rotation 90",
			direction: utils.Vector2{X: 1, Y: 0},
			vertices: [3]*utils.Vector2{
				{X: 130, Y: 100},
				{X: 70, Y: 120},
				{X: 70, Y: 80},
			},
		},
		{
			name:      "Rotation 180",
			direction: utils.Vector2{X: 0, Y: 1},
			vertices: [3]*utils.Vector2{
				{X: 100, Y: 130},
				{X: 80, Y: 70},
				{X: 120, Y: 70},
			},
		},
		{
			name:      "Rotation 270",
			direction: utils.Vector2{X: -1, Y: 0},
			vertices: [3]*utils.Vector2{
				{X: 70, Y: 100},
				{X: 130, Y: 80},
				{X: 130, Y: 120},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p.Direction = c.direction
			assert.Equal(c.vertices, p.Triangle())
		})
	}

	t.Run("Rotation 45", func(t *testing.T) {
		// for drawing, y-axis is inverted. Thus the expected values are inverted
		p.Direction = *utils.NewVector2(1, 0).Rotate(-45)
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

func TestPlayerGetHitboxCircle(t *testing.T) {
	assert := assert.New(t)

	center := utils.Vector2{X: 100, Y: 100}
	radius := 20
	p := &sprite.Player{
		Circle: sprite.Circle{
			Center: center,
			Radius: radius,
		},
	}

	pos, rad := p.GetHitboxCircule()

	assert.Equal(center, pos)
	assert.Equal(radius, rad)
}

func TestPlayerIsCollided(t *testing.T) {
	assert := assert.New(t)

	p := &sprite.Player{
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
			collided := p.IsCollided(c.hitbox)
			assert.Equal(c.collided, collided)
		})
	}
}

func TestPlayerFire(t *testing.T) {
	assert := assert.New(t)
	gunConfig := sprite.GunConfig{Radius: 5, Speed: 10.0, RateLimit: time.Millisecond * 500}
	p := &sprite.Player{
		Circle: sprite.Circle{
			Center: utils.Vector2{X: 100, Y: 100},
			Radius: 20,
		},
		Gun: gunConfig,
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
