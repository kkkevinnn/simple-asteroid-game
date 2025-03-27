package sprite

import (
	"image"
	"log"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AsteroidControl struct {
	AsteroidFactory   *AsteroidFactory
	Asteroids         []*Asteroid
	AsteroidRadiusMin int
	AsteroidKind      int
	Bounds            image.Rectangle
	SpawnRate         time.Duration
	lastSpwan         time.Time
}

func NewAsteroidControl(radiusMin int, kind int, bounds image.Rectangle, spwanRate string) *AsteroidControl {
	dur, err := time.ParseDuration(spwanRate)
	if err != nil {
		log.Fatal(err)
	}
	return &AsteroidControl{
		AsteroidFactory:   NewAsteroidFactory(radiusMin, kind, bounds, 100, 40, 30),
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

	// remove bullets that are out of bounds
	for _, a := range c.Asteroids {
		if a.Center.X < -float64(a.Radius) || a.Center.X > float64(c.Bounds.Max.X+a.Radius) ||
			a.Center.Y < -float64(a.Radius) || a.Center.Y > float64(c.Bounds.Max.Y+a.Radius) {
			a.Destory()
		}
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
	edge := rand.IntN(4)
	return c.AsteroidFactory.NewAsteroid(edge)
}

func randIntRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func (c *AsteroidControl) HitAsteroid(i int) {
	if i >= len(c.Asteroids) || c.Asteroids[i].IsDestoryed() {
		return
	}

	if c.Asteroids[i].Radius > c.AsteroidRadiusMin {
		newRadius := c.Asteroids[i].Radius - c.AsteroidRadiusMin
		newSpeed := c.Asteroids[i].Speed * 1.2
		newAngel := rand.Float64()*30 + 20
		newDirection1 := c.Asteroids[i].Direction.Clone().Rotate(newAngel)
		newDirection2 := c.Asteroids[i].Direction.Clone().Rotate(-newAngel)
		newCenter1 := c.Asteroids[i].Center.Clone().Add(*newDirection1.Clone().Scale(float64(c.Asteroids[i].Radius)))
		newCenter2 := c.Asteroids[i].Center.Clone().Add(*newDirection2.Clone().Scale(float64(c.Asteroids[i].Radius)))

		c.AddAsteroid(NewAsteroid(*newCenter1, newRadius, newSpeed, *newDirection1))
		c.AddAsteroid(NewAsteroid(*newCenter2, newRadius, newSpeed, *newDirection2))

		log.Printf("2 Asteroids created:\n")
		log.Printf("1. %v\n", c.Asteroids[len(c.Asteroids)-2])
		log.Printf("2. %v\n", c.Asteroids[len(c.Asteroids)-1])
	}
	c.Asteroids[i].Destory()
}

func (c *AsteroidControl) Clean() {
	mark := len(c.Asteroids)
	for i := len(c.Asteroids) - 1; i >= 0; i-- {
		if c.Asteroids[i].IsDestoryed() {
			mark--
			tmpAsteroid := c.Asteroids[i]
			c.Asteroids[i] = c.Asteroids[mark]
			c.Asteroids[mark] = tmpAsteroid
		}
	}

	if cleaned := len(c.Asteroids) - mark; cleaned > 0 {
		log.Printf("Cleaned %d asteroids out of %d\n", cleaned, len(c.Asteroids))
	}
	c.Asteroids = c.Asteroids[:mark]
}
