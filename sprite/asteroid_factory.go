package sprite

import (
	"image"
	"math/rand/v2"

	"asteroid/utils"
)

type AsteroidFactory struct {
	MinRadius int
	Kind      int
	Bounds    image.Rectangle
	MaxSpeed  float64
	MinSpeed  float64
	MaxAngle  float64
}

var factory *AsteroidFactory

func NewAsteroidFactory(minRadius int, kind int, bounds image.Rectangle, maxSpeed float64, minSpeed float64, maxAngle float64) *AsteroidFactory {
	if factory == nil {
		factory = &AsteroidFactory{
			MinRadius: minRadius,
			Kind:      kind,
			Bounds:    bounds,
			MaxSpeed:  maxSpeed,
			MinSpeed:  minSpeed,
			MaxAngle:  maxAngle,
		}
	}
	return factory
}

func (af *AsteroidFactory) NewAsteroid(edge int) *Asteroid {
	radius := (rand.IntN(af.Kind) + 1) * af.MinRadius
	speed := float64(randIntRange(int(af.MinSpeed), int(af.MaxSpeed)))
	angle := rand.NormFloat64() * af.MaxAngle

	var center utils.Vector2
	var direction *utils.Vector2

	switch edge {
	case 0:
		center = utils.Vector2{X: float64(randIntRange(radius, af.Bounds.Max.X-radius)), Y: 0}
		direction = utils.NewVector2(0, 1).Rotate(angle)
	case 1:
		center = utils.Vector2{X: float64(af.Bounds.Max.X - radius), Y: float64(randIntRange(radius, af.Bounds.Max.Y-radius))}
		direction = utils.NewVector2(-1, 0).Rotate(angle)
	case 2:
		center = utils.Vector2{X: float64(randIntRange(radius, af.Bounds.Max.X-radius)), Y: float64(af.Bounds.Max.Y - radius)}
		direction = utils.NewVector2(0, -1).Rotate(angle)
	case 3:
		center = utils.Vector2{X: float64(af.Bounds.Min.X + radius), Y: float64(randIntRange(radius, af.Bounds.Max.Y-radius))}
		direction = utils.NewVector2(1, 0).Rotate(angle)
	}

	return NewAsteroid(center, radius, speed, *direction)
}
