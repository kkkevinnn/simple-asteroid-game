package sprite_test

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"

	"asteroid/sprite"
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

func TestBulletControlHitBullet_Clean(t *testing.T) {
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	bulletControl := sprite.NewBulletControl(bounds)

	bulletControl.AddBullet(&sprite.Bullet{})
	bulletControl.AddBullet(&sprite.Bullet{})

	assert := assert.New(t)
	bulletControl.HitBullet(0)
	assert.Equal(2, len(bulletControl.Bullets))
	assert.Equal(true, bulletControl.Bullets[0].IsDestoryed())

	bulletControl.Clean()
	assert.Equal(1, len(bulletControl.Bullets))
}
