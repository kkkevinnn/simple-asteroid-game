package sprite

import (
	"asteroid/utils"
	"image"
	"log"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AsteroidControl struct {
	Asteroids         []*Asteroid
	AsteroidRadiusMax int
	AsteroidRadiusMin int
	AsteroidKind      int
	Bounds            image.Rectangle
	SpawnRate         time.Duration
	lastSpwan         time.Time
}

func NewAsteroidControl(radiusMin int, radiusMax int, kind int, bounds image.Rectangle, spwanRate string) *AsteroidControl {
	dur, err := time.ParseDuration(spwanRate)
	if err != nil {
		log.Fatal(err)
	}
	return &AsteroidControl{
		AsteroidRadiusMax: radiusMax,
		AsteroidRadiusMin: radiusMin,
		AsteroidKind:      kind,
		Bounds:            bounds,
		SpawnRate:         dur,
	}
}

func (c *AsteroidControl) Update() {
	for _, a := range c.Asteroids {
		a.Update()
	}

	if time.Since(c.lastSpwan) >= c.SpawnRate {
		c.lastSpwan = time.Now()
		c.AddAsteroid(c.SpawnAsteroid())
	}
}

func (c *AsteroidControl) Draw(screen *ebiten.Image) {
	for _, a := range c.Asteroids {
		a.Draw(screen)
	}
}

func (c *AsteroidControl) AddAsteroid(a *Asteroid) {
	c.Asteroids = append(c.Asteroids, a)
}

func (c *AsteroidControl) SpawnAsteroid() *Asteroid {
	// radius := rand.IntN(c.AsteroidRadiusMax+1-c.AsteroidRadiusMin) + c.AsteroidRadiusMin
	radius := rand.IntN(c.AsteroidKind) * c.AsteroidRadiusMin
	speed := float64(randIntRange(40, 100))
	angle := rand.NormFloat64() * 30

	edge := rand.IntN(4)
	var center utils.Vector2
	var direction *utils.Vector2

	switch edge {
	case 0:
		center = utils.Vector2{X: float64(randIntRange(radius, c.Bounds.Max.X-radius)), Y: 0}
		direction = utils.NewVector2(0, 1).Rotate(angle)
	case 1:
		center = utils.Vector2{X: float64(c.Bounds.Max.X - radius), Y: float64(randIntRange(radius, c.Bounds.Max.Y-radius))}
		direction = utils.NewVector2(-1, 0).Rotate(angle)
	case 2:
		center = utils.Vector2{X: float64(randIntRange(radius, c.Bounds.Max.X-radius)), Y: float64(c.Bounds.Max.Y - radius)}
		direction = utils.NewVector2(0, -1).Rotate(angle)
	case 3:
		center = utils.Vector2{X: float64(c.Bounds.Min.X + radius), Y: float64(randIntRange(radius, c.Bounds.Max.Y-radius))}
		direction = utils.NewVector2(1, 0).Rotate(angle)
	}

	return NewAsteroid(center, radius, speed, *direction)
}

func randIntRange(min, max int) int {
	return rand.IntN(max-min) + min
}
