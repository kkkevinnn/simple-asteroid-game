package sprite_test

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"

	"asteroid/sprite"
	"asteroid/utils"
)

func TestNewBulletControl(t *testing.T) {
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	bulletControl := sprite.NewBulletControl(bounds)

	assert := assert.New(t)
	assert.NotNil(bulletControl)
	assert.Equal(bounds, bulletControl.Bounds)
	assert.Equal(0, len(bulletControl.Bullets))
}

func TestBulletControlAddBullet(t *testing.T) {
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	bulletControl := sprite.NewBulletControl(bounds)

	bullet := &sprite.Bullet{}
	bulletControl.AddBullet(bullet)

	assert := assert.New(t)
	assert.Equal(1, len(bulletControl.Bullets))
	assert.Equal(bullet, bulletControl.Bullets[0])
}

func TestBulletControlHitBullet(t *testing.T) {
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	bulletControl := sprite.NewBulletControl(bounds)

	bulletControl.AddBullet(&sprite.Bullet{})
	bulletControl.AddBullet(&sprite.Bullet{})

	assert := assert.New(t)
	bulletControl.HitBullet(0)
	assert.Equal(2, len(bulletControl.Bullets))
	assert.Equal(true, bulletControl.Bullets[0].IsDestoryed())
}

func TestBullectControlClean(t *testing.T) {
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	bc := sprite.NewBulletControl(bounds)

	bc.AddBullet(&sprite.Bullet{})
	bc.AddBullet(&sprite.Bullet{})

	assert := assert.New(t)
	bc.Clean()
	assert.Equal(2, len(bc.Bullets))
	bc.Bullets[0].Destory()
	bc.Clean()
	assert.Equal(1, len(bc.Bullets))
}

func TestBulletControlUpdate(t *testing.T) {
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	bc := sprite.NewBulletControl(bounds)

	oldCenter := utils.Vector2{X: 100, Y: 100}
	bc.AddBullet(&sprite.Bullet{
		Circle: sprite.Circle{
			Center:    oldCenter,
			Radius:    10,
			Speed:     400,
			Direction: utils.Vector2{X: 0, Y: -1},
		},
	})
	bc.AddBullet(&sprite.Bullet{
		Circle: sprite.Circle{
			Center: utils.Vector2{X: -100, Y: -100},
		},
	})

	assert := assert.New(t)
	bc.Update()
	assert.NotEqual(oldCenter, bc.Bullets[0].Circle.Center)
	assert.Equal(true, bc.Bullets[1].IsDestoryed())
}
