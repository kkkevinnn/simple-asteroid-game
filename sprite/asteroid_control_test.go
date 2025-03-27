package sprite_test

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"

	"asteroid/sprite"
	"asteroid/utils"
)

func TestNewAsteroidControl(t *testing.T) {
	minRadius := 20
	kind := 3
	bounds := image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}
	spawnRate := "1s"

	asteroidControl := sprite.NewAsteroidControl(minRadius, kind, bounds, spawnRate)
	assert := assert.New(t)
	assert.NotNil(asteroidControl)
	assert.Equal(minRadius, asteroidControl.AsteroidRadiusMin)
	assert.Equal(kind, asteroidControl.AsteroidKind)
	assert.Equal(bounds, asteroidControl.Bounds)
	assert.Equal(spawnRate, asteroidControl.SpawnRate.String())
}

func TestAsteroidControlUpdate(t *testing.T) {
	ac := sprite.NewAsteroidControl(20, 3, image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}, "1s")
	ac.Update()
	assert := assert.New(t)
	assert.Equal(1, len(ac.Asteroids))

	ac.AddAsteroid(&sprite.Asteroid{Circle: sprite.Circle{Center: utils.Vector2{X: -100, Y: -100}}})
	assert.Equal(2, len(ac.Asteroids))
	ac.Update()
	assert.Equal(2, len(ac.Asteroids))
	assert.Equal(true, ac.Asteroids[1].IsDestoryed())
}

func TestAddAsteroid(t *testing.T) {
	asteroidControl := sprite.NewAsteroidControl(20, 3, image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}, "1s")
	asteroidControl.AddAsteroid(&sprite.Asteroid{})

	assert.Equal(t, 1, len(asteroidControl.Asteroids))
}

func TestAsteroidControlHitAsteroid(t *testing.T) {
	asteroidControl := sprite.NewAsteroidControl(20, 3, image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}, "1s")
	asteroidControl.AddAsteroid(&sprite.Asteroid{Circle: sprite.Circle{Radius: 40}})
	asteroidControl.AddAsteroid(&sprite.Asteroid{Circle: sprite.Circle{Radius: 20}})
	asteroidControl.AddAsteroid(&sprite.Asteroid{Circle: sprite.Circle{Radius: 20}})

	assert.Equal(t, 3, len(asteroidControl.Asteroids))
	asteroidControl.HitAsteroid(0)
	assert.Equal(t, true, asteroidControl.Asteroids[0].IsDestoryed())
	assert.Equal(t, 5, len(asteroidControl.Asteroids))

	asteroidControl.HitAsteroid(1)
	assert.Equal(t, true, asteroidControl.Asteroids[1].IsDestoryed())
	assert.Equal(t, 5, len(asteroidControl.Asteroids))

	asteroidControl.Clean()
	assert.Equal(t, 3, len(asteroidControl.Asteroids))
}

func TestAsteroidControlClean(t *testing.T) {
	ac := sprite.NewAsteroidControl(20, 3, image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}, "1s")
	ac.AddAsteroid(&sprite.Asteroid{})
	ac.AddAsteroid(&sprite.Asteroid{})

	assert := assert.New(t)
	ac.Clean()
	assert.Equal(2, len(ac.Asteroids))
	ac.Asteroids[0].Destory()
	ac.Clean()
	assert.Equal(1, len(ac.Asteroids))
}

func TestAsteroidControlSpawnAsteroid(t *testing.T) {
	asteroidControl := sprite.NewAsteroidControl(20, 3, image.Rectangle{Max: image.Point{X: 1000, Y: 1000}}, "1s")
	asteroid := asteroidControl.SpawnAsteroid()

	assert := assert.New(t)
	assert.NotNil(asteroid)
}
